// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/skip-mev/slinky/x/marketmap/types"
)

// Oracle is an autogenerated mock type for the Oracle type
type Oracle struct {
	mock.Mock
}

// GetLastSyncTime provides a mock function with given fields:
func (_m *Oracle) GetLastSyncTime() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLastSyncTime")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// GetMarketMap provides a mock function with given fields:
func (_m *Oracle) GetMarketMap() *types.MarketMap {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMarketMap")
	}

	var r0 *types.MarketMap
	if rf, ok := ret.Get(0).(func() *types.MarketMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MarketMap)
		}
	}

	return r0
}

// GetPrices provides a mock function with given fields:
func (_m *Oracle) GetPrices() map[string]*big.Float {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPrices")
	}

	var r0 map[string]*big.Float
	if rf, ok := ret.Get(0).(func() map[string]*big.Float); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Float)
		}
	}

	return r0
}

// IsRunning provides a mock function with given fields:
func (_m *Oracle) IsRunning() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRunning")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Oracle) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Oracle) Stop() {
	_m.Called()
}

// NewOracle creates a new instance of Oracle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracle(t interface {
	mock.TestingT
	Cleanup(func())
}) *Oracle {
	mock := &Oracle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
