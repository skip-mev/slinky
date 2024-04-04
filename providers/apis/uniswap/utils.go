package uniswap

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/constants"
	"github.com/skip-mev/slinky/oracle/types"
)

// Name is the name of the Uniswap API.
const Name = "uniswap_api"

// PoolConfig is the configuration for a Uniswap V3 pool. This is specific to each pair of tokens.
type PoolConfig struct {
	// Address is the Uniswap V3 pool address.
	Address string `json:"address"`
	// BaseDecimals is the number of decimals for the base token. This should be derived from the
	// token contract.
	BaseDecimals int `json:"base_decimals"`
	// QuoteDecimals is the number of decimals for the quote token. This should be derived from the
	// token contract.
	QuoteDecimals int `json:"quote_decimals"`
	// Invert is utilized to invert the price of a pool's reserves. This may be required for certain
	// pools as the price is derived based on the sorted order of the ERC20 addresses of the tokens
	// in the pool.
	Invert bool `json:"invert"`
}

// ValidateBasic validates the pool configuration.
func (pc *PoolConfig) ValidateBasic() error {
	if !common.IsHexAddress(pc.Address) {
		return fmt.Errorf("pool address is not a valid ethereum address")
	}

	if pc.BaseDecimals <= 0 {
		return fmt.Errorf("base decimals must be positive")
	}

	if pc.QuoteDecimals <= 0 {
		return fmt.Errorf("quote decimals must be positive")
	}

	return nil
}

var (
	// DefaultAPIConfig is the default configuration for the Uniswap API.
	DefaultAPIConfig = config.APIConfig{
		Name:             Name,
		Atomic:           false,
		Enabled:          true,
		Timeout:          1000 * time.Millisecond,
		Interval:         2000 * time.Millisecond,
		ReconnectTimeout: 2000 * time.Millisecond,
		MaxQueries:       1,
		URL:              "https://eth.public-rpc.com/",
	}

	// DefaultMarketConfig is the default market configuration for Uniswap.
	DefaultMarketConfig = types.TickerToProviderConfig{
		constants.WETH_USDC: {
			Name:           Name,
			OffChainTicker: "WETH/USDC",
		},
	}
)
