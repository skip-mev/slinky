// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// BankKeeper is an autogenerated mock type for the BankKeeper type
type BankKeeper struct {
	mock.Mock
}

// BurnCoins provides a mock function with given fields: ctx, moduleName, amounts
func (_m *BankKeeper) BurnCoins(ctx context.Context, moduleName string, amounts types.Coins) error {
	ret := _m.Called(ctx, moduleName, amounts)

	if len(ret) == 0 {
		panic("no return value specified for BurnCoins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, moduleName, amounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MintCoins provides a mock function with given fields: ctx, moduleName, amt
func (_m *BankKeeper) MintCoins(ctx context.Context, moduleName string, amt types.Coins) error {
	ret := _m.Called(ctx, moduleName, amt)

	if len(ret) == 0 {
		panic("no return value specified for MintCoins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, moduleName, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoinsFromAccountToModule provides a mock function with given fields: ctx, senderAddr, recipientModule, amt
func (_m *BankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	ret := _m.Called(ctx, senderAddr, recipientModule, amt)

	if len(ret) == 0 {
		panic("no return value specified for SendCoinsFromAccountToModule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, string, types.Coins) error); ok {
		r0 = rf(ctx, senderAddr, recipientModule, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoinsFromModuleToAccount provides a mock function with given fields: ctx, senderModule, recipientAddr, amt
func (_m *BankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, senderModule, recipientAddr, amt)

	if len(ret) == 0 {
		panic("no return value specified for SendCoinsFromModuleToAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, senderModule, recipientAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBankKeeper creates a new instance of BankKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBankKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *BankKeeper {
	mock := &BankKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
