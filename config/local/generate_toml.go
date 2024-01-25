//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/providers/websockets/kucoin"

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

var oracleCfgPath = flag.String("oracle-config-path", "oracle.toml", "path to write the oracle config file to")

// LocalConfig defines a readable config for local development. Any changes to this
// file should be reflected in oracle.toml. To update the oracle.toml file, run
// `make update-local-config`. This will update any changes to the oracle.toml file
// as they are made to this file.
var LocalConfig = config.OracleConfig{
	// -----------------------------------------------------------	//
	// --------------------All Currency Pairs---------------------	//
	// -----------------------------------------------------------	//
	CurrencyPairs: []oracletypes.CurrencyPair{
		oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		oracletypes.NewCurrencyPair("ATOM", "USD"),
		oracletypes.NewCurrencyPair("SOLANA", "USD"),
		oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		oracletypes.NewCurrencyPair("AVAX", "USD"),
		oracletypes.NewCurrencyPair("DYDX", "USD"),
		oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		oracletypes.NewCurrencyPair("OSMOSIS", "USD"),
	},
	Production: false,
	// -----------------------------------------------------------	//
	// ----------------------Metrics Config-----------------------	//
	// -----------------------------------------------------------	//
	Metrics: config.MetricsConfig{
		Enabled:                 true,
		PrometheusServerAddress: "localhost:8000",
	},
	UpdateInterval: 1 * time.Second,
	Providers: []config.ProviderConfig{
		// -----------------------------------------------------------	//
		// ---------------------Start API Providers--------------------	//
		// -----------------------------------------------------------	//
		//
		// NOTE: Some of the provider's are only capable of fetching data for a subset of
		// all of the currency pairs. Before adding a new market to the oracle, ensure that
		// the provider supports fetching data for the currency pair.
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Binance API----------------------	//
		// 	Name: binance.Name,
		// 	API:  binance.DefaultUSAPIConfig,
		// 	Market: config.MarketConfig{
		// 		Name: binance.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "BTCUSDT",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETHUSDT",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOMUSDT",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOLUSDT",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETHBTC",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Coinbase API--------------------	//
		// 	Name: coinbaseapi.Name,
		// 	API:  coinbaseapi.DefaultAPIConfig,
		// 	Market: config.MarketConfig{
		// 		Name: coinbaseapi.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "BTC-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETH-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOM-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOL-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "TIA-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"AVAX/USD": {
		// 				Ticker:       "AVAX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "DYDX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETH-BTC",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 			"OSMOSIS/USD": {
		// 				Ticker:       "OSMO-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("OSMOSIS", "USD"),
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Coingecko API-------------------	//
		// 	Name: coingecko.Name,
		// 	API:  coingecko.DefaultAPIConfig,
		// 	Market: config.MarketConfig{
		// 		Name: coingecko.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "bitcoin/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ethereum/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "cosmos/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "solana/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "celestia/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "dydx-chain/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ethereum/btc",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 			"OSMOSIS/USD": {
		// 				Ticker:       "osmosis/usd",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("OSMOSIS", "USD"),
		// 			},
		// 		},
		// 	},
		// },
		// // -----------------------------------------------------------	//
		// // ---------------------Start WebSocket Providers--------------	//
		// // -----------------------------------------------------------	//
		// //
		// // NOTE: Some of the provider's are only capable of fetching data for a subset of
		// // all of the currency pairs. Before adding a new market to the oracle, ensure that
		// // the provider supports fetching data for the currency pair.
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Coinbase WebSocket--------------	//
		// 	Name:      coinbasews.Name,
		// 	WebSocket: coinbasews.DefaultWebSocketConfig,
		// 	Market: config.MarketConfig{
		// 		Name: coinbasews.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "BTC-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETH-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOM-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOL-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "TIA-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"AVAX/USD": {
		// 				Ticker:       "AVAX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "DYDX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETH-BTC",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 			"OSMOSIS/USD": {
		// 				Ticker:       "OSMO-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("OSMOSIS", "USD"),
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Crypto.com WebSocket------------	//
		// 	Name:      cryptodotcom.Name,
		// 	WebSocket: cryptodotcom.DefaultWebSocketConfig,
		// 	Market: config.MarketConfig{
		// 		Name: cryptodotcom.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "BTCUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETHUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOMUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOLUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "TIAUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"AVAX/USD": {
		// 				Ticker:       "AVAXUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "DYDXUSD-PERP",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETH_BTC",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 			"OSMOSIS/USD": {
		// 				Ticker:       "OSMO_USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("OSMOSIS", "USD"),
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start Kraken WebSocket----------------	//
		// 	Name:      kraken.Name,
		// 	WebSocket: kraken.DefaultWebSocketConfig,
		// 	Market: config.MarketConfig{
		// 		Name: kraken.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "XBT/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETH/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOM/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOL/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "TIA/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"AVAX/USD": {
		// 				Ticker:       "AVAX/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "DYDX/USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETH/XBT",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 		},
		// 	},
		// },
		{
			// -----------------------------------------------------------	//
			// ---------------------Start Kucoin WebSocket----------------	//
			Name:      kucoin.Name,
			WebSocket: kucoin.DefaultWebSocketConfig,
			Market: config.MarketConfig{
				Name: kucoin.Name,
				CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
					"BITCOIN/USD": {
						Ticker:       "BTC-USDT",
						CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
					},
				},
			},
		},
		// {
		// 	// -----------------------------------------------------------	//
		// 	// ---------------------Start OKX WebSocket-------------------	//
		// 	Name:      okx.Name,
		// 	WebSocket: okx.DefaultWebSocketConfig,
		// 	Market: config.MarketConfig{
		// 		Name: okx.Name,
		// 		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
		// 			"BITCOIN/USD": {
		// 				Ticker:       "BTC-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD"),
		// 			},
		// 			"ETHEREUM/USD": {
		// 				Ticker:       "ETH-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD"),
		// 			},
		// 			"ATOM/USD": {
		// 				Ticker:       "ATOM-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ATOM", "USD"),
		// 			},
		// 			"SOLANA/USD": {
		// 				Ticker:       "SOL-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD"),
		// 			},
		// 			"CELESTIA/USD": {
		// 				Ticker:       "TIA-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("CELESTIA", "USD"),
		// 			},
		// 			"AVAX/USD": {
		// 				Ticker:       "AVAX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD"),
		// 			},
		// 			"DYDX/USD": {
		// 				Ticker:       "DYDX-USD",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD"),
		// 			},
		// 			"ETHEREUM/BITCOIN": {
		// 				Ticker:       "ETH-BTC",
		// 				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
		// 			},
		// 		},
		// 	},
		// },
	},
}

// main executes a simple script that encodes the local config file to the local
// directory.
func main() {
	flag.Parse()

	// Open the local config file. This will overwrite any changes made to the
	// local config file.
	f, err := os.Create(*oracleCfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating local config file: %v\n", err)
	}
	defer f.Close()

	// Encode the local config file.
	encoder := toml.NewEncoder(f)
	if err := encoder.Encode(LocalConfig); err != nil {
		panic(err)
	}
}
