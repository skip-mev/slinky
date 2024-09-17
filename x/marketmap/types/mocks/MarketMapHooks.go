// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MarketMapHooks is an autogenerated mock type for the MarketMapHooks type
type MarketMapHooks struct {
	mock.Mock
}

type MarketMapHooks_Expecter struct {
	mock *mock.Mock
}

func (_m *MarketMapHooks) EXPECT() *MarketMapHooks_Expecter {
	return &MarketMapHooks_Expecter{mock: &_m.Mock}
}

// AfterMarketCreated provides a mock function with given fields: ctx, market
func (_m *MarketMapHooks) AfterMarketCreated(ctx types.Context, market marketmaptypes.Market) error {
	ret := _m.Called(ctx, market)

	if len(ret) == 0 {
		panic("no return value specified for AfterMarketCreated")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, marketmaptypes.Market) error); ok {
		r0 = rf(ctx, market)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarketMapHooks_AfterMarketCreated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterMarketCreated'
type MarketMapHooks_AfterMarketCreated_Call struct {
	*mock.Call
}

// AfterMarketCreated is a helper method to define mock.On call
//   - ctx types.Context
//   - market marketmaptypes.Market
func (_e *MarketMapHooks_Expecter) AfterMarketCreated(ctx interface{}, market interface{}) *MarketMapHooks_AfterMarketCreated_Call {
	return &MarketMapHooks_AfterMarketCreated_Call{Call: _e.mock.On("AfterMarketCreated", ctx, market)}
}

func (_c *MarketMapHooks_AfterMarketCreated_Call) Run(run func(ctx types.Context, market marketmaptypes.Market)) *MarketMapHooks_AfterMarketCreated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(marketmaptypes.Market))
	})
	return _c
}

func (_c *MarketMapHooks_AfterMarketCreated_Call) Return(_a0 error) *MarketMapHooks_AfterMarketCreated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MarketMapHooks_AfterMarketCreated_Call) RunAndReturn(run func(types.Context, marketmaptypes.Market) error) *MarketMapHooks_AfterMarketCreated_Call {
	_c.Call.Return(run)
	return _c
}

// AfterMarketGenesis provides a mock function with given fields: ctx, tickers
func (_m *MarketMapHooks) AfterMarketGenesis(ctx types.Context, tickers map[string]marketmaptypes.Market) error {
	ret := _m.Called(ctx, tickers)

	if len(ret) == 0 {
		panic("no return value specified for AfterMarketGenesis")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, map[string]marketmaptypes.Market) error); ok {
		r0 = rf(ctx, tickers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarketMapHooks_AfterMarketGenesis_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterMarketGenesis'
type MarketMapHooks_AfterMarketGenesis_Call struct {
	*mock.Call
}

// AfterMarketGenesis is a helper method to define mock.On call
//   - ctx types.Context
//   - tickers map[string]marketmaptypes.Market
func (_e *MarketMapHooks_Expecter) AfterMarketGenesis(ctx interface{}, tickers interface{}) *MarketMapHooks_AfterMarketGenesis_Call {
	return &MarketMapHooks_AfterMarketGenesis_Call{Call: _e.mock.On("AfterMarketGenesis", ctx, tickers)}
}

func (_c *MarketMapHooks_AfterMarketGenesis_Call) Run(run func(ctx types.Context, tickers map[string]marketmaptypes.Market)) *MarketMapHooks_AfterMarketGenesis_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(map[string]marketmaptypes.Market))
	})
	return _c
}

func (_c *MarketMapHooks_AfterMarketGenesis_Call) Return(_a0 error) *MarketMapHooks_AfterMarketGenesis_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MarketMapHooks_AfterMarketGenesis_Call) RunAndReturn(run func(types.Context, map[string]marketmaptypes.Market) error) *MarketMapHooks_AfterMarketGenesis_Call {
	_c.Call.Return(run)
	return _c
}

// AfterMarketUpdated provides a mock function with given fields: ctx, market
func (_m *MarketMapHooks) AfterMarketUpdated(ctx types.Context, market marketmaptypes.Market) error {
	ret := _m.Called(ctx, market)

	if len(ret) == 0 {
		panic("no return value specified for AfterMarketUpdated")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, marketmaptypes.Market) error); ok {
		r0 = rf(ctx, market)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarketMapHooks_AfterMarketUpdated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterMarketUpdated'
type MarketMapHooks_AfterMarketUpdated_Call struct {
	*mock.Call
}

// AfterMarketUpdated is a helper method to define mock.On call
//   - ctx types.Context
//   - market marketmaptypes.Market
func (_e *MarketMapHooks_Expecter) AfterMarketUpdated(ctx interface{}, market interface{}) *MarketMapHooks_AfterMarketUpdated_Call {
	return &MarketMapHooks_AfterMarketUpdated_Call{Call: _e.mock.On("AfterMarketUpdated", ctx, market)}
}

func (_c *MarketMapHooks_AfterMarketUpdated_Call) Run(run func(ctx types.Context, market marketmaptypes.Market)) *MarketMapHooks_AfterMarketUpdated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(marketmaptypes.Market))
	})
	return _c
}

func (_c *MarketMapHooks_AfterMarketUpdated_Call) Return(_a0 error) *MarketMapHooks_AfterMarketUpdated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MarketMapHooks_AfterMarketUpdated_Call) RunAndReturn(run func(types.Context, marketmaptypes.Market) error) *MarketMapHooks_AfterMarketUpdated_Call {
	_c.Call.Return(run)
	return _c
}

// NewMarketMapHooks creates a new instance of MarketMapHooks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMarketMapHooks(t interface {
	mock.TestingT
	Cleanup(func())
}) *MarketMapHooks {
	mock := &MarketMapHooks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
