// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorStore is an autogenerated mock type for the ValidatorStore type
type ValidatorStore struct {
	mock.Mock
}

// TotalBondedTokens provides a mock function with given fields: ctx
func (_m *ValidatorStore) TotalBondedTokens(ctx context.Context) (math.Int, error) {
	ret := _m.Called(ctx)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (math.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) math.Int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorByConsAddr provides a mock function with given fields: ctx, addr
func (_m *ValidatorStore) ValidatorByConsAddr(ctx context.Context, addr types.ConsAddress) (stakingtypes.ValidatorI, error) {
	ret := _m.Called(ctx, addr)

	var r0 stakingtypes.ValidatorI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ConsAddress) (stakingtypes.ValidatorI, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ConsAddress) stakingtypes.ValidatorI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.ValidatorI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ConsAddress) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewValidatorStore creates a new instance of ValidatorStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidatorStore(t interface {
	mock.TestingT
	Cleanup(func())
},
) *ValidatorStore {
	mock := &ValidatorStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
