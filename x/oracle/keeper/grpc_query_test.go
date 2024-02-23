package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/mock"

	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	"github.com/skip-mev/slinky/x/oracle/keeper"
	"github.com/skip-mev/slinky/x/oracle/types"
)

func (s *KeeperTestSuite) TestGetAllCurrencyPairs() {
	qs := keeper.NewQueryServer(s.oracleKeeper)

	// test that an error is returned if no CurrencyPairs have been registered in the module
	s.Run("an error is returned if no CurrencyPairs have been registered in the module", func() {
		// execute query
		_, err := qs.GetAllCurrencyPairs(s.ctx, nil)
		s.Require().NotNil(s.T(), err)
	})

	// test that after CurrencyPairs are registered, all of them are returned from the query
	s.Run("after CurrencyPairs are registered, all of them are returned from the query", func() {
		// insert multiple currency Pairs
		s.Require().NoError(s.oracleKeeper.CreateCurrencyPair(s.ctx, slinkytypes.CurrencyPair{
			Base:  "AA",
			Quote: "BB",
		}))
		s.Require().NoError(s.oracleKeeper.CreateCurrencyPair(s.ctx, slinkytypes.CurrencyPair{
			Base:  "CC",
			Quote: "DD",
		}))
		s.Require().NoError(s.oracleKeeper.CreateCurrencyPair(s.ctx, slinkytypes.CurrencyPair{
			Base:  "EE",
			Quote: "FF",
		}))

		// manually insert a new CurrencyPair as well
		s.Require().NoError(s.oracleKeeper.SetPriceForCurrencyPair(s.ctx, slinkytypes.CurrencyPair{
			Base:  "GG",
			Quote: "HH",
		}, types.QuotePrice{Price: sdkmath.NewInt(100)}))

		expectedCurrencyPairs := map[string]struct{}{"AA/BB": {}, "CC/DD": {}, "EE/FF": {}, "GG/HH": {}}

		// query for pairs
		res, err := qs.GetAllCurrencyPairs(s.ctx, nil)
		s.Require().Nil(err)

		// assert that currency-pairs are correctly returned
		for _, cp := range res.CurrencyPairs {
			_, ok := expectedCurrencyPairs[cp.String()]
			s.Require().True(ok)
		}
	})
}

func (s *KeeperTestSuite) TestGetPrice() {
	// set CPs on genesis for testing
	cpg := []types.CurrencyPairGenesis{
		{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "AA",
				Quote: "ETHEREUM",
			},
			CurrencyPairPrice: &types.QuotePrice{
				Price: sdkmath.NewInt(100),
			},
			Nonce: 12,
			Id:    2,
		},
		{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "CC",
				Quote: "BB",
			},
			Id: 1,
		},
	}

	// init genesis
	s.oracleKeeper.InitGenesis(s.ctx, *types.NewGenesisState(cpg, 3))

	tcs := []struct {
		name       string
		setup      func()
		req        *types.GetPriceRequest
		res        *types.GetPriceResponse
		expectPass bool
	}{
		{
			"if the request is nil, expect failure - fail",
			func() {},
			nil,
			nil,
			false,
		},
		{
			"if the currency pair selector is nil, expect failure - fail",
			func() {},
			&types.GetPriceRequest{
				CurrencyPairSelector: nil,
			},
			nil,
			false,
		},
		{
			"if the currency pair selector's currency pair is nil, expect failure - fail",
			func() {},
			&types.GetPriceRequest{
				CurrencyPairSelector: &types.GetPriceRequest_CurrencyPair{CurrencyPair: nil},
			},
			nil,
			false,
		},
		{
			"if the query is for a currency pair that does not exist fail - fail",
			func() {},
			&types.GetPriceRequest{
				CurrencyPairSelector: &types.GetPriceRequest_CurrencyPairId{CurrencyPairId: "DD/EE"},
			},
			nil,
			false,
		},
		{
			"if the query is for a currency-pair with no price, only the nonce (0) is returned - pass",
			func() {
				s.mockMarketMapKeeper.On("GetTicker", mock.Anything, mock.Anything).Return(marketmaptypes.Ticker{
					CurrencyPair:     slinkytypes.CurrencyPair{Base: "CC", Quote: "BB"},
					Decimals:         8,
					MinProviderCount: 3,
					Metadata_JSON:    "",
				}, nil).Once()
			},
			&types.GetPriceRequest{
				CurrencyPairSelector: &types.GetPriceRequest_CurrencyPairId{CurrencyPairId: "CC/BB"},
			},
			&types.GetPriceResponse{
				Nonce:    0,
				Decimals: uint64(8),
				Id:       1,
			},
			true,
		},
		{
			"if the query is for a currency pair that has valid price data, return the price + the nonce - pass",
			func() {
				s.mockMarketMapKeeper.On("GetTicker", mock.Anything, mock.Anything).Return(marketmaptypes.Ticker{
					CurrencyPair:     slinkytypes.CurrencyPair{Base: "AA", Quote: "ETHEREUM"},
					Decimals:         18,
					MinProviderCount: 3,
					Metadata_JSON:    "",
				}, nil).Once()
			},
			&types.GetPriceRequest{
				CurrencyPairSelector: &types.GetPriceRequest_CurrencyPair{CurrencyPair: &slinkytypes.CurrencyPair{Base: "AA", Quote: "ETHEREUM"}},
			},
			&types.GetPriceResponse{
				Nonce: 12,
				Price: &types.QuotePrice{
					Price: sdkmath.NewInt(100),
				},
				Decimals: uint64(18),
				Id:       2,
			},
			true,
		},
	}

	qs := keeper.NewQueryServer(s.oracleKeeper)

	for _, tc := range tcs {
		s.Run(tc.name, func() {
			tc.setup()

			// get the response + error from the query
			res, err := qs.GetPrice(s.ctx, tc.req)
			if !tc.expectPass {
				s.Require().NotNil(err)
				return
			}

			// otherwise, assert no error, and check response
			s.Require().Nil(err)

			// check response
			s.Require().Equal(res.Nonce, tc.res.Nonce)

			// check price if possible
			if tc.res.Price != nil {
				checkQuotePriceEqual(s.T(), *tc.res.Price, *res.Price)
			}

			// check decimals
			s.Require().Equal(tc.res.Decimals, res.Decimals)

			// check id
			s.Require().Equal(tc.res.Id, res.Id)
		})
	}
}
