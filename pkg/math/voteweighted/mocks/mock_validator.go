// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	math "cosmossdk.io/math"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/crypto/types"
)

// ValidatorI is an autogenerated mock type for the ValidatorI type
type ValidatorI struct {
	mock.Mock
}

// ConsPubKey provides a mock function with given fields:
func (_m *ValidatorI) ConsPubKey() (types.PubKey, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConsPubKey")
	}

	var r0 types.PubKey
	var r1 error
	if rf, ok := ret.Get(0).(func() (types.PubKey, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() types.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PubKey)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBondedTokens provides a mock function with given fields:
func (_m *ValidatorI) GetBondedTokens() math.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBondedTokens")
	}

	var r0 math.Int
	if rf, ok := ret.Get(0).(func() math.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// GetCommission provides a mock function with given fields:
func (_m *ValidatorI) GetCommission() math.LegacyDec {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCommission")
	}

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func() math.LegacyDec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// GetConsAddr provides a mock function with given fields:
func (_m *ValidatorI) GetConsAddr() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConsAddr")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsensusPower provides a mock function with given fields: _a0
func (_m *ValidatorI) GetConsensusPower(_a0 math.Int) int64 {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetConsensusPower")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func(math.Int) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetDelegatorShares provides a mock function with given fields:
func (_m *ValidatorI) GetDelegatorShares() math.LegacyDec {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDelegatorShares")
	}

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func() math.LegacyDec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// GetMinSelfDelegation provides a mock function with given fields:
func (_m *ValidatorI) GetMinSelfDelegation() math.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMinSelfDelegation")
	}

	var r0 math.Int
	if rf, ok := ret.Get(0).(func() math.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// GetMoniker provides a mock function with given fields:
func (_m *ValidatorI) GetMoniker() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMoniker")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOperator provides a mock function with given fields:
func (_m *ValidatorI) GetOperator() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOperator")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStatus provides a mock function with given fields:
func (_m *ValidatorI) GetStatus() stakingtypes.BondStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 stakingtypes.BondStatus
	if rf, ok := ret.Get(0).(func() stakingtypes.BondStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(stakingtypes.BondStatus)
	}

	return r0
}

// GetTokens provides a mock function with given fields:
func (_m *ValidatorI) GetTokens() math.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTokens")
	}

	var r0 math.Int
	if rf, ok := ret.Get(0).(func() math.Int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// IsBonded provides a mock function with given fields:
func (_m *ValidatorI) IsBonded() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsBonded")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsJailed provides a mock function with given fields:
func (_m *ValidatorI) IsJailed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsJailed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUnbonded provides a mock function with given fields:
func (_m *ValidatorI) IsUnbonded() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsUnbonded")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUnbonding provides a mock function with given fields:
func (_m *ValidatorI) IsUnbonding() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsUnbonding")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SharesFromTokens provides a mock function with given fields: amt
func (_m *ValidatorI) SharesFromTokens(amt math.Int) (math.LegacyDec, error) {
	ret := _m.Called(amt)

	if len(ret) == 0 {
		panic("no return value specified for SharesFromTokens")
	}

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(math.Int) (math.LegacyDec, error)); ok {
		return rf(amt)
	}
	if rf, ok := ret.Get(0).(func(math.Int) math.LegacyDec); ok {
		r0 = rf(amt)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(math.Int) error); ok {
		r1 = rf(amt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SharesFromTokensTruncated provides a mock function with given fields: amt
func (_m *ValidatorI) SharesFromTokensTruncated(amt math.Int) (math.LegacyDec, error) {
	ret := _m.Called(amt)

	if len(ret) == 0 {
		panic("no return value specified for SharesFromTokensTruncated")
	}

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(math.Int) (math.LegacyDec, error)); ok {
		return rf(amt)
	}
	if rf, ok := ret.Get(0).(func(math.Int) math.LegacyDec); ok {
		r0 = rf(amt)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(math.Int) error); ok {
		r1 = rf(amt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TmConsPublicKey provides a mock function with given fields:
func (_m *ValidatorI) TmConsPublicKey() (crypto.PublicKey, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TmConsPublicKey")
	}

	var r0 crypto.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func() (crypto.PublicKey, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() crypto.PublicKey); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(crypto.PublicKey)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokensFromShares provides a mock function with given fields: _a0
func (_m *ValidatorI) TokensFromShares(_a0 math.LegacyDec) math.LegacyDec {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for TokensFromShares")
	}

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func(math.LegacyDec) math.LegacyDec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// TokensFromSharesRoundUp provides a mock function with given fields: _a0
func (_m *ValidatorI) TokensFromSharesRoundUp(_a0 math.LegacyDec) math.LegacyDec {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for TokensFromSharesRoundUp")
	}

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func(math.LegacyDec) math.LegacyDec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// TokensFromSharesTruncated provides a mock function with given fields: _a0
func (_m *ValidatorI) TokensFromSharesTruncated(_a0 math.LegacyDec) math.LegacyDec {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for TokensFromSharesTruncated")
	}

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func(math.LegacyDec) math.LegacyDec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// NewValidatorI creates a new instance of ValidatorI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidatorI(t interface {
	mock.TestingT
	Cleanup(func())
}) *ValidatorI {
	mock := &ValidatorI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
