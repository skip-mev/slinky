package oracle

import (
	"context"
	"fmt"
	"sync"
	"time"

	"cosmossdk.io/log"
	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/service/servers/oracle/types"
	"google.golang.org/grpc"
)

var _ OracleClient = (*PriceDaemon)(nil)

type PriceDaemon struct {
	logger log.Logger

	// config is the configuration of the daemon.
	config config.AppConfig
	// client is the underlying oracle client used to fetch prices.
	OracleClient
	// latestResponse is the latest price response fetched by the daemon.
	resp ThreadSafeResponse
	// doneCh is a channel that is closed when the daemon is stopped.
	doneCh chan struct{}
}

// NewPriceDaemon creates a new price daemon with the given configuration.
func NewPriceDaemon(
	logger log.Logger,
	cfg config.AppConfig,
	client OracleClient,
) (*PriceDaemon, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}

	if err := cfg.ValidateBasic(); err != nil {
		return nil, err
	}

	if client == nil {
		return nil, fmt.Errorf("oracle client cannot be nil")
	}

	return &PriceDaemon{
		logger:       logger.With("process", "price_daemon"),
		config:       cfg,
		OracleClient: client,
		doneCh:       make(chan struct{}),
	}, nil
}

// Start starts the price daemon. This method will block until the daemon is stopped.
func (d *PriceDaemon) Start(ctx context.Context) error {
	if err := d.OracleClient.Start(ctx); err != nil {
		return err
	}
	defer d.OracleClient.Stop()

	ticker := time.NewTicker(d.config.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			d.logger.Info("stopping price daemon from context")
			return ctx.Err()
		case <-d.doneCh:
			d.logger.Info("price daemon stopped")
			close(d.doneCh)
			return nil
		case <-ticker.C:
			d.fetchPrices(ctx)
		}
	}
}

// fetchPrices fetches the latest prices from the oracle client.
func (d *PriceDaemon) fetchPrices(ctx context.Context) {
	d.logger.Debug("fetching prices")

	fetchCtx, cancel := context.WithTimeout(ctx, d.config.ClientTimeout)
	defer cancel()

	resp, err := d.OracleClient.Prices(fetchCtx, &types.QueryPricesRequest{})
	if err != nil {
		d.logger.Error(
			"failed to fetch prices from sidecar",
			"err", err,
			"address", d.config.OracleAddress,
		)

		return
	}

	ts := time.Now()
	d.logger.Debug("fetched prices", "timestamp", ts, "prices", resp.Prices)
	d.resp.Update(resp)
}

// Prices returns the latest price response fetched by the daemon. If the latest response
// is too stale, an error is returned.
func (d *PriceDaemon) Prices(
	_ context.Context,
	_ *types.QueryPricesRequest,
	_ ...grpc.CallOption,
) (*types.QueryPricesResponse, error) {
	latest, ts := d.resp.Get()
	if latest == nil {
		return nil, fmt.Errorf("no prices fetched by price daemon yet")
	}

	if time.Since(ts) > d.config.MaxAge {
		return nil, fmt.Errorf(
			"latest prices from the price daemon are too stale; last fetched at %s; diff %s ago",
			ts.Format(time.RFC3339),
			time.Since(ts).String(),
		)
	}

	return latest, nil
}

// Stop stops the price daemon.
func (d *PriceDaemon) Stop() error {
	d.doneCh <- struct{}{}
	return nil
}

// ThreadSafeResponse is a thread-safe wrapper around a QueryPricesResponse.
type ThreadSafeResponse struct {
	sync.Mutex

	resp      *types.QueryPricesResponse
	timestamp time.Time
}

// NewThreadSafeResponse creates a new thread-safe response with the given response.
func NewThreadSafeResponse() *ThreadSafeResponse {
	return &ThreadSafeResponse{
		resp:      nil,
		timestamp: time.Time{},
	}
}

// Update updates the response and timestamp of the thread-safe response.
func (r *ThreadSafeResponse) Update(resp *types.QueryPricesResponse) {
	r.Lock()
	defer r.Unlock()

	r.resp = resp
	r.timestamp = time.Now()
}

// Get returns the response and timestamp of the thread-safe response.
func (r *ThreadSafeResponse) Get() (*types.QueryPricesResponse, time.Time) {
	r.Lock()
	defer r.Unlock()

	return r.resp, r.timestamp
}
