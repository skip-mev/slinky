// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	keeper "github.com/skip-mev/connect/v2/x/sla/keeper"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Keeper is an autogenerated mock type for the Keeper type
type Keeper struct {
	mock.Mock
}

type Keeper_Expecter struct {
	mock *mock.Mock
}

func (_m *Keeper) EXPECT() *Keeper_Expecter {
	return &Keeper_Expecter{mock: &_m.Mock}
}

// UpdatePriceFeeds provides a mock function with given fields: ctx, updates
func (_m *Keeper) UpdatePriceFeeds(ctx types.Context, updates keeper.PriceFeedUpdates) error {
	ret := _m.Called(ctx, updates)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePriceFeeds")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, keeper.PriceFeedUpdates) error); ok {
		r0 = rf(ctx, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Keeper_UpdatePriceFeeds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePriceFeeds'
type Keeper_UpdatePriceFeeds_Call struct {
	*mock.Call
}

// UpdatePriceFeeds is a helper method to define mock.On call
//   - ctx types.Context
//   - updates keeper.PriceFeedUpdates
func (_e *Keeper_Expecter) UpdatePriceFeeds(ctx interface{}, updates interface{}) *Keeper_UpdatePriceFeeds_Call {
	return &Keeper_UpdatePriceFeeds_Call{Call: _e.mock.On("UpdatePriceFeeds", ctx, updates)}
}

func (_c *Keeper_UpdatePriceFeeds_Call) Run(run func(ctx types.Context, updates keeper.PriceFeedUpdates)) *Keeper_UpdatePriceFeeds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(keeper.PriceFeedUpdates))
	})
	return _c
}

func (_c *Keeper_UpdatePriceFeeds_Call) Return(_a0 error) *Keeper_UpdatePriceFeeds_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Keeper_UpdatePriceFeeds_Call) RunAndReturn(run func(types.Context, keeper.PriceFeedUpdates) error) *Keeper_UpdatePriceFeeds_Call {
	_c.Call.Return(run)
	return _c
}

// NewKeeper creates a new instance of Keeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *Keeper {
	mock := &Keeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
