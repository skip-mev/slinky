// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	pkgtypes "github.com/skip-mev/slinky/pkg/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CurrencyPairStrategy is an autogenerated mock type for the CurrencyPairStrategy type
type CurrencyPairStrategy struct {
	mock.Mock
}

// FromID provides a mock function with given fields: ctx, id
func (_m *CurrencyPairStrategy) FromID(ctx types.Context, id uint64) (pkgtypes.CurrencyPair, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FromID")
	}

	var r0 pkgtypes.CurrencyPair
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (pkgtypes.CurrencyPair, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) pkgtypes.CurrencyPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pkgtypes.CurrencyPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDecodedPrice provides a mock function with given fields: ctx, cp, priceBytes
func (_m *CurrencyPairStrategy) GetDecodedPrice(ctx types.Context, cp pkgtypes.CurrencyPair, priceBytes []byte) (*big.Int, error) {
	ret := _m.Called(ctx, cp, priceBytes)

	if len(ret) == 0 {
		panic("no return value specified for GetDecodedPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair, []byte) (*big.Int, error)); ok {
		return rf(ctx, cp, priceBytes)
	}
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair, []byte) *big.Int); ok {
		r0 = rf(ctx, cp, priceBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, pkgtypes.CurrencyPair, []byte) error); ok {
		r1 = rf(ctx, cp, priceBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEncodedPrice provides a mock function with given fields: ctx, cp, price
func (_m *CurrencyPairStrategy) GetEncodedPrice(ctx types.Context, cp pkgtypes.CurrencyPair, price *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, cp, price)

	if len(ret) == 0 {
		panic("no return value specified for GetEncodedPrice")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair, *big.Int) ([]byte, error)); ok {
		return rf(ctx, cp, price)
	}
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair, *big.Int) []byte); ok {
		r0 = rf(ctx, cp, price)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, pkgtypes.CurrencyPair, *big.Int) error); ok {
		r1 = rf(ctx, cp, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxNumCP provides a mock function with given fields: ctx
func (_m *CurrencyPairStrategy) GetMaxNumCP(ctx types.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMaxNumCP")
	}

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

// ID provides a mock function with given fields: ctx, cp
func (_m *CurrencyPairStrategy) ID(ctx types.Context, cp pkgtypes.CurrencyPair) (uint64, error) {
	ret := _m.Called(ctx, cp)

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) (uint64, error)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair) uint64); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context, pkgtypes.CurrencyPair) error); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCurrencyPairStrategy creates a new instance of CurrencyPairStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCurrencyPairStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *CurrencyPairStrategy {
	mock := &CurrencyPairStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
