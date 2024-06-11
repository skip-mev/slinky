package oracle

import (
	"go.uber.org/zap"

	oraclemetrics "github.com/skip-mev/slinky/oracle/metrics"
	"github.com/skip-mev/slinky/oracle/types"
	mmclienttypes "github.com/skip-mev/slinky/service/clients/marketmap/types"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
)

// Option is a functional option for the market map state.
type Option func(impl *OracleImpl)

// WithLogger sets the logger for the provider orchestrator.
func WithLogger(logger *zap.Logger) Option {
	return func(m *OracleImpl) {
		if logger == nil {
			panic("logger cannot be nil")
		}

		m.logger = logger.With(zap.String("process", "oracle"))
	}
}

// WithMarketMap sets the market map for the provider orchestrator.
func WithMarketMap(marketMap mmtypes.MarketMap) Option {
	return func(m *OracleImpl) {
		if err := marketMap.ValidateBasic(); err != nil {
			panic(err)
		}

		m.marketMap = marketMap
	}
}

// WithPriceAPIQueryHandlerFactory sets the Price API query handler factory for the provider orchestrator.
// Specifically, this is what is utilized to construct price providers that are API based.
func WithPriceAPIQueryHandlerFactory(factory types.PriceAPIQueryHandlerFactory) Option {
	return func(m *OracleImpl) {
		if factory == nil {
			panic("api query handler factory cannot be nil")
		}

		m.priceAPIFactory = factory
	}
}

// WithPriceWebSocketQueryHandlerFactory sets the websocket query handler factory for the provider orchestrator.
// Specifically, this is what is utilized to construct price providers that are websocket based.
func WithPriceWebSocketQueryHandlerFactory(factory types.PriceWebSocketQueryHandlerFactory) Option {
	return func(m *OracleImpl) {
		if factory == nil {
			panic("websocket query handler factory cannot be nil")
		}

		m.priceWSFactory = factory
	}
}

// WithMarketMapperFactory sets the market map factory for the provider orchestrator.
// Specifically, this is what is utilized to construct market map providers.
func WithMarketMapperFactory(factory mmclienttypes.MarketMapFactory) Option {
	return func(m *OracleImpl) {
		if factory == nil {
			panic("market map factory cannot be nil")
		}

		m.marketMapperFactory = factory
	}
}

// WithWriteTo sets the file path to which market map updates will be written to. Note that this is optional.
func WithWriteTo(filePath string) Option {
	return func(m *OracleImpl) {
		m.writeTo = filePath
	}
}

func WithPriceProviders(pps ...*types.PriceProvider) Option {
	return func(m *OracleImpl) {
		for _, pp := range pps {
			m.priceProviders[pp.Name()] = ProviderState{Provider: pp}
		}
	}
}

func WithMetrics(met oraclemetrics.Metrics) Option {
	return func(m *OracleImpl) {
		m.metrics = met
	}
}

func WithMarketMapProvider(mmp *mmclienttypes.MarketMapProvider) Option {
	return func(impl *OracleImpl) {
		impl.mmProvider = mmp
	}
}
