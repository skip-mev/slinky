package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/skip-mev/slinky/oracle/types"
	pkgtypes "github.com/skip-mev/slinky/pkg/types"
	mmtypes "github.com/skip-mev/slinky/x/mm2/types"
)

func TestProviderTickersFromMarketMap(t *testing.T) {
	cases := []struct {
		name     string
		provider string
		market   mmtypes.MarketMap
		expected []types.ProviderTicker
		err      bool
	}{
		{
			name:     "empty market map",
			provider: "test",
			market:   mmtypes.MarketMap{},
			expected: nil,
			err:      false,
		},
		{
			name:     "invalid market map",
			provider: "test",
			market: mmtypes.MarketMap{
				Markets: map[string]mmtypes.Market{
					"test": {},
				},
			},
			expected: nil,
			err:      true,
		},
		{
			name:     "single market for the provider",
			provider: "test",
			market: mmtypes.MarketMap{
				Markets: map[string]mmtypes.Market{
					"BTC/USD": {
						Ticker: mmtypes.NewTicker("BTC", "USD", 8, 1),
						ProviderConfigs: []mmtypes.ProviderConfig{
							{
								Name:           "test",
								OffChainTicker: "BTC/USDT",
								Metadata_JSON:  "{}",
							},
						},
					},
				},
			},
			expected: []types.ProviderTicker{
				types.NewProviderTicker(
					"test",
					"BTC/USDT",
					"{}",
					types.DefaultTickerDecimals,
				),
			},
			err: false,
		},
		{
			name:     "provider has no configs in market map",
			provider: "test",
			market: mmtypes.MarketMap{
				Markets: map[string]mmtypes.Market{
					"BTC/USD": {
						Ticker: mmtypes.NewTicker("BTC", "USD", 8, 1),
						ProviderConfigs: []mmtypes.ProviderConfig{
							{
								Name:           "other",
								OffChainTicker: "BTC/USDT",
								Metadata_JSON:  "{}",
							},
						},
					},
				},
			},
			expected: nil,
			err:      false,
		},
		{
			name:     "duplicate markets for the provider",
			provider: "test",
			market: mmtypes.MarketMap{
				Markets: map[string]mmtypes.Market{
					"ETH/USD": {
						Ticker: mmtypes.NewTicker("ETH", "USD", 8, 1),
						ProviderConfigs: []mmtypes.ProviderConfig{
							{
								Name:           "test",
								OffChainTicker: "ETH/USDT",
								NormalizeByPair: &pkgtypes.CurrencyPair{
									Base:  "USDT",
									Quote: "USD",
								},
								Metadata_JSON: "{}",
							},
						},
					},
					"USDT/USD": {
						Ticker: mmtypes.NewTicker("USDT", "USD", 8, 1),
						ProviderConfigs: []mmtypes.ProviderConfig{
							{
								Name:           "test",
								OffChainTicker: "ETH/USDT",
								Invert:         true,
								NormalizeByPair: &pkgtypes.CurrencyPair{
									Base:  "ETH",
									Quote: "USD",
								},
								Metadata_JSON: "{}",
							},
						},
					},
				},
			},
			expected: []types.ProviderTicker{
				types.NewProviderTicker(
					"test",
					"ETH/USDT",
					"{}",
					types.DefaultTickerDecimals,
				),
			},
			err: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := types.ProviderTickersFromMarketMap(tc.provider, tc.market)
			if tc.err {
				require.Error(t, err)
				return
			}

			expectedCache := make(map[types.ProviderTicker]struct{})
			for _, ticker := range tc.expected {
				expectedCache[ticker] = struct{}{}
			}
			actualCache := make(map[types.ProviderTicker]struct{})
			for _, ticker := range actual {
				actualCache[ticker] = struct{}{}
			}
			require.Equal(t, expectedCache, actualCache)
		})
	}
}

func TestSimpleProviderTicker(t *testing.T) {
	cases := []struct {
		name   string
		ticker types.SimpleProviderTicker
		err    bool
	}{
		{
			name: "valid simple provider ticker",
			ticker: types.SimpleProviderTicker{
				Name:           "test",
				OffChainTicker: "BTC/USD",
				JSON:           "{}",
			},
			err: false,
		},
		{
			name: "empty name",
			ticker: types.SimpleProviderTicker{
				Name:           "",
				OffChainTicker: "BTC/USD",
				JSON:           "{}",
			},
			err: true,
		},
		{
			name: "empty off-chain ticker",
			ticker: types.SimpleProviderTicker{
				Name:           "test",
				OffChainTicker: "",
				JSON:           "{}",
			},
			err: true,
		},
		{
			name: "empty JSON",
			ticker: types.SimpleProviderTicker{
				Name:           "test",
				OffChainTicker: "BTC/USD",
				JSON:           "",
			},
			err: false,
		},
		{

			name: "invalid JSON",
			ticker: types.SimpleProviderTicker{
				Name:           "test",
				OffChainTicker: "BTC/USD",
				JSON:           "{",
			},
			err: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.ticker.ValidateBasic()
			if tc.err {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
