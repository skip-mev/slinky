package keeper_test

import (
	sdkmath "cosmossdk.io/math"

	"github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

func (s *KeeperTestSuite) TestInitGenesisInvalidGenesis() {
	s.Run("test that init genesis with invalid genesis params - fails", func() {
		// create a fake genesis state with invalid params
		gs := types.GenesisState{
			MarketMap: types.DefaultGenesisState().MarketMap,
			Params: types.Params{
				MarketAuthority: "invalid",
				Version:         0,
			},
		}

		// assert that InitGenesis panics
		s.Require().Panics(func() {
			s.keeper.InitGenesis(s.ctx, gs)
		})
	})

	s.Run("test that init genesis with invalid duplicate runs - fails", func() {
		// create a valid genesis
		gs := types.DefaultGenesisState()

		gs.MarketMap = types.MarketMap{
			Tickers: map[string]types.Ticker{
				ethusdt.String(): ethusdt,
				btcusdt.String(): btcusdt,
				usdcusd.String(): usdcusd,
			},
			Paths: map[string]types.Paths{
				ethusdt.String(): ethusdtPaths,
				btcusdt.String(): btcusdtPaths,
				usdcusd.String(): usdcusdPaths,
			},
			Providers: map[string]types.Providers{
				ethusdt.String(): ethusdtProviders,
				btcusdt.String(): btcusdtProviders,
				usdcusd.String(): usdcusdProviders,
			},
		}

		// assert that InitGenesis panics
		s.Require().Panics(func() {
			s.keeper.InitGenesis(s.ctx, *gs)
			s.keeper.InitGenesis(s.ctx, *gs)
		})
	})
}

func (s *KeeperTestSuite) TestInitGenesisValid() {
	s.Run("init valid default genesis", func() {
		gs := types.DefaultGenesisState()

		s.Require().NotPanics(func() {
			s.keeper.InitGenesis(s.ctx, *gs)
		})
	})

	s.Run("init valid genesis with fields", func() {
		// first register x/oracle genesis
		ogs := oracletypes.DefaultGenesisState()
		ogs.NextId = 3
		ogs.CurrencyPairGenesis = []oracletypes.CurrencyPairGenesis{
			{
				CurrencyPair:      ethusdt.CurrencyPair,
				CurrencyPairPrice: &oracletypes.QuotePrice{Price: sdkmath.NewInt(19)},
				Nonce:             0,
				Id:                0,
			},
			{
				CurrencyPair:      btcusdt.CurrencyPair,
				CurrencyPairPrice: &oracletypes.QuotePrice{Price: sdkmath.NewInt(19)},
				Nonce:             0,
				Id:                1,
			},
			{
				CurrencyPair:      usdcusd.CurrencyPair,
				CurrencyPairPrice: nil,
				Nonce:             0,
				Id:                2,
			},
		}

		s.Require().NotPanics(func() {
			s.oracleKeeper.InitGenesis(s.ctx, *ogs)
		})

		gs := types.DefaultGenesisState()
		gs.MarketMap = types.MarketMap{
			Tickers: map[string]types.Ticker{
				ethusdt.String(): ethusdt,
				btcusdt.String(): btcusdt,
				usdcusd.String(): usdcusd,
			},
			Paths: map[string]types.Paths{
				ethusdt.String(): ethusdtPaths,
				btcusdt.String(): btcusdtPaths,
				usdcusd.String(): usdcusdPaths,
			},
			Providers: map[string]types.Providers{
				ethusdt.String(): ethusdtProviders,
				btcusdt.String(): btcusdtProviders,
				usdcusd.String(): usdcusdProviders,
			},
		}

		s.Require().NotPanics(func() {
			s.keeper.InitGenesis(s.ctx, *gs)
		})

		gotTickers, err := s.keeper.GetAllTickersMap(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(gs.MarketMap.Tickers, gotTickers)

		gotPaths, err := s.keeper.GetAllPathsMap(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(gs.MarketMap.Paths, gotPaths)

		gotProviders, err := s.keeper.GetAllProvidersMap(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(gs.MarketMap.Providers, gotProviders)
	})
}
