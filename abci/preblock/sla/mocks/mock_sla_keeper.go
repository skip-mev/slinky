// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	keeper "github.com/skip-mev/slinky/x/sla/keeper"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Keeper is an autogenerated mock type for the Keeper type
type Keeper struct {
	mock.Mock
}

// UpdatePriceFeeds provides a mock function with given fields: ctx, updates
func (_m *Keeper) UpdatePriceFeeds(ctx types.Context, updates keeper.PriceFeedUpdates) error {
	ret := _m.Called(ctx, updates)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, keeper.PriceFeedUpdates) error); ok {
		r0 = rf(ctx, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKeeper creates a new instance of Keeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Keeper {
	mock := &Keeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
