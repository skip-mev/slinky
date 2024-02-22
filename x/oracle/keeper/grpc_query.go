package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	"github.com/skip-mev/slinky/x/oracle/types"
)

// queryServer is the default implementation of the x/oracle QueryService.
type queryServer struct {
	k Keeper
}

// NewQueryServer returns an implementation of the x/oracle QueryServer.
func NewQueryServer(k Keeper) types.QueryServer {
	return queryServer{
		k,
	}
}

var _ types.QueryServer = queryServer{}

// GetAllCurrencyPairs returns the set of all currency pairs that the module is tracking QuotePrices for.
// It returns an error to the caller if there are no CurrencyPairs being tracked by the module.
func (q queryServer) GetAllCurrencyPairs(ctx context.Context, _ *types.GetAllCurrencyPairsRequest) (*types.GetAllCurrencyPairsResponse, error) {
	// get all currency pairs from state
	cps := q.k.GetAllCurrencyPairs(sdk.UnwrapSDKContext(ctx))

	// if no currency pairs exist in the module, return an error to indicate to caller
	if len(cps) == 0 {
		return &types.GetAllCurrencyPairsResponse{}, nil
	}

	return &types.GetAllCurrencyPairsResponse{
		CurrencyPairs: cps,
	}, nil
}

// GetPrice gets the QuotePrice and the nonce for the QuotePrice for a given CurrencyPair. The request contains a
// CurrencyPairSelector (either the stringified CurrencyPair, or the CurrencyPair itself). If the request is nil this method fails.
// If the selector is an incorrectly formatted string this method fails. If the QuotePrice / Nonce do not exist for this CurrencyPair, this method fails.
func (q queryServer) GetPrice(goCtx context.Context, req *types.GetPriceRequest) (_ *types.GetPriceResponse, err error) {
	var cp slinkytypes.CurrencyPair

	// fail on nil requests
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	// determine what type the selector from the response is giving
	switch cpI := req.CurrencyPairSelector.(type) {

	case *types.GetPriceRequest_CurrencyPairId:
		// retrieve the currency pair from the stringified ID, and fail if incorrectly formatted
		cp, err = slinkytypes.CurrencyPairFromString(cpI.CurrencyPairId)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling CurrencyPairID: %w", err)
		}

	case *types.GetPriceRequest_CurrencyPair:
		// retrieve CurrencyPair directly from selector
		if cpI.CurrencyPair == nil {
			return nil, fmt.Errorf("currency Pair cannot be nil")
		}
		cp = *cpI.CurrencyPair

	default:
		// fail if any other type of CurrencyPairSelector is given
		return nil, fmt.Errorf("invalid CurrencyPairSelector given in request (consult documentation)")
	}

	// unwrap ctx
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the QuotePrice + nonce for the given CurrencyPair
	qpn, err := q.k.GetPriceWithNonceForCurrencyPair(ctx, cp)
	if err != nil {
		return nil, fmt.Errorf("no price / nonce reported for CurrencyPair: %s, the module is not tracking this CurrencyPair", cp.String())
	}

	id, ok := q.k.GetIDForCurrencyPair(ctx, cp)
	if !ok {
		return nil, fmt.Errorf("no ID found for CurrencyPair: %s", cp.String())
	}

	ticker, err := k.mmKeeper.GetTicker(ctx, cp.String())
	if err != nil {
		return nil, fmt.Errorf("no ticker for CurrencyPair")
	}

	// return the QuotePrice + Nonce
	return &types.GetPriceResponse{
		Price:    &qpn.QuotePrice,
		Nonce:    qpn.Nonce(),
		Decimals: uint64(cp.Decimals()),
		Id:       id,
	}, nil
}

// GetPrices gets the array of the QuotePrice and the nonce for the QuotePrice for a given CurrencyPairs.
func (q queryServer) GetPrices(goCtx context.Context, req *types.GetPricesRequest) (_ *types.GetPricesResponse, err error) {
	var cp slinkytypes.CurrencyPair

	// fail on nil requests
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	prices := make([]types.GetPriceResponse, 0, len(req.CurrencyPairIds))
	for _, cid := range req.CurrencyPairIds {
		cp, err = slinkytypes.CurrencyPairFromString(cid)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling CurrencyPairID: %w", err)
		}

		// unwrap ctx
		ctx := sdk.UnwrapSDKContext(goCtx)

		// get the QuotePrice + nonce for the given CurrencyPair
		qpn, err := q.k.GetPriceWithNonceForCurrencyPair(ctx, cp)
		if err != nil {
			return nil, fmt.Errorf("no price / nonce reported for CurrencyPair: %v, the module is not tracking this CurrencyPair", cp)
		}

		id, ok := q.k.GetIDForCurrencyPair(ctx, cp)
		if !ok {
			return nil, fmt.Errorf("no ID found for CurrencyPair: %v", cp)
		}

		prices = append(prices, types.GetPriceResponse{
			Price:    &qpn.QuotePrice,
			Nonce:    qpn.Nonce(),
			Decimals: uint64(cp.Decimals()),
			Id:       id,
		})
	}

	return &types.GetPricesResponse{
		Prices: prices,
	}, nil
}
