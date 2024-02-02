package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgAddCurrencyPairs{}
	_ sdk.Msg = &MsgRemoveCurrencyPairs{}
)

// NewMsgAddCurrencyPairs returns a new message from a set of currency-pairs and an authority
func NewMsgAddCurrencyPairs(authority string, cps []CurrencyPair) MsgAddCurrencyPairs {
	return MsgAddCurrencyPairs{
		Authority:     authority,
		CurrencyPairs: cps,
	}
}

// GetSigners gets the addresses that must sign this message. In this case, the signer
// must be the module authority.
func (m MsgAddCurrencyPairs) GetSigners() []sdk.AccAddress {
	// convert from string to acc address
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic determines whether the information in the message is formatted correctly, specifically
// whether the authority is a valid acc-address, and that each CurrencyPair in the message is formatted correctly
func (m MsgAddCurrencyPairs) ValidateBasic() error {
	// validate authority address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	// validate currency pairs
	for _, cp := range m.CurrencyPairs {
		if err := cp.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// NewMsgRemoveCurrencyPairs returns a new message to remove a set of currency-pairs from the x/oracle module's state
func NewMsgRemoveCurrencyPairs(authority string, currencyPairIds []string) MsgRemoveCurrencyPairs {
	return MsgRemoveCurrencyPairs{
		Authority:       authority,
		CurrencyPairIds: currencyPairIds,
	}
}

// GetSigners gets the addresses that must sign this message. In this case, the signer
// must be the module authority.
func (m MsgRemoveCurrencyPairs) GetSigners() []sdk.AccAddress {
	// convert from string to acc address
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic determines whether the information in the message is valid, specifically
// whether the authority is a valid acc-address, and that each CurrencyPairID in the message is formatted correctly
func (m MsgRemoveCurrencyPairs) ValidateBasic() error {
	// validate authority address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	// check that each CurrencyPairID is correctly formatted
	for _, id := range m.CurrencyPairIds {
		if _, err := CurrencyPairFromID(id, 1); err != nil {
			return err
		}
	}

	return nil
}
