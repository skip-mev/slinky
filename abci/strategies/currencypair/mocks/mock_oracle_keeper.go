// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	mock "github.com/stretchr/testify/mock"

	pkgtypes "github.com/skip-mev/slinky/pkg/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// OracleKeeper is an autogenerated mock type for the OracleKeeper type
type OracleKeeper struct {
	mock.Mock
}

// GetAllCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetAllCurrencyPairs(ctx types.Context) []pkgtypes.CurrencyPair {
	ret := _m.Called(ctx)

	var r0 []pkgtypes.CurrencyPair
	if rf, ok := ret.Get(0).(func(types.Context) []pkgtypes.CurrencyPair); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pkgtypes.CurrencyPair)
		}
	}

	return r0
}

// GetCurrencyPairFromID provides a mock function with given fields: ctx, id
func (_m *OracleKeeper) GetCurrencyPairFromID(ctx types.Context, id uint64) (pkgtypes.CurrencyPair, bool) {
	ret := _m.Called(ctx, id)

	var r0 pkgtypes.CurrencyPair
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (pkgtypes.CurrencyPair, bool)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) pkgtypes.CurrencyPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pkgtypes.CurrencyPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetIDForCurrencyPair provides a mock function with given fields: ctx, cp
func (_m *OracleKeeper) GetIDForCurrencyPair(ctx types.Context, cp pkgtypes.CurrencyPair) (uint64, bool) {
	ret := _m.Called(ctx, cp)

	var r0 uint64
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) (uint64, bool)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) uint64); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context, pkgtypes.CurrencyPair) bool); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNumCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetNumCurrencyPairs(ctx types.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumRemovedCurrencyPairs provides a mock function with given fields: ctx
func (_m *OracleKeeper) GetNumRemovedCurrencyPairs(ctx types.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPriceForCurrencyPair provides a mock function with given fields: ctx, cp
func (_m *OracleKeeper) GetPriceForCurrencyPair(ctx types.Context, cp pkgtypes.CurrencyPair) (oracletypes.QuotePrice, error) {
	ret := _m.Called(ctx, cp)

	var r0 oracletypes.QuotePrice
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) (oracletypes.QuotePrice, error)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) oracletypes.QuotePrice); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(oracletypes.QuotePrice)
	}

	if rf, ok := ret.Get(1).(func(types.Context, pkgtypes.CurrencyPair) error); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOracleKeeper creates a new instance of OracleKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracleKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *OracleKeeper {
	mock := &OracleKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
