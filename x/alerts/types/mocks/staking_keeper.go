// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	math "cosmossdk.io/math"
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// BondDenom provides a mock function with given fields: ctx
func (_m *StakingKeeper) BondDenom(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BondDenom")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatorByConsAddr provides a mock function with given fields: ctx, consAddr
func (_m *StakingKeeper) GetValidatorByConsAddr(ctx context.Context, consAddr cosmos_sdktypes.ConsAddress) (stakingtypes.Validator, error) {
	ret := _m.Called(ctx, consAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetValidatorByConsAddr")
	}

	var r0 stakingtypes.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress) (stakingtypes.Validator, error)); ok {
		return rf(ctx, consAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress) stakingtypes.Validator); ok {
		r0 = rf(ctx, consAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.ConsAddress) error); ok {
		r1 = rf(ctx, consAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Slash provides a mock function with given fields: ctx, consAddr, infractionHeight, power, slashFactor
func (_m *StakingKeeper) Slash(ctx context.Context, consAddr cosmos_sdktypes.ConsAddress, infractionHeight int64, power int64, slashFactor math.LegacyDec) (math.Int, error) {
	ret := _m.Called(ctx, consAddr, infractionHeight, power, slashFactor)

	if len(ret) == 0 {
		panic("no return value specified for Slash")
	}

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress, int64, int64, math.LegacyDec) (math.Int, error)); ok {
		return rf(ctx, consAddr, infractionHeight, power, slashFactor)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress, int64, int64, math.LegacyDec) math.Int); ok {
		r0 = rf(ctx, consAddr, infractionHeight, power, slashFactor)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.ConsAddress, int64, int64, math.LegacyDec) error); ok {
		r1 = rf(ctx, consAddr, infractionHeight, power, slashFactor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStakingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
