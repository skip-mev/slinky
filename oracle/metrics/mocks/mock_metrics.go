// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

type Metrics_Expecter struct {
	mock *mock.Mock
}

func (_m *Metrics) EXPECT() *Metrics_Expecter {
	return &Metrics_Expecter{mock: &_m.Mock}
}

// AddProviderCountForMarket provides a mock function with given fields: pairID, count
func (_m *Metrics) AddProviderCountForMarket(pairID string, count int) {
	_m.Called(pairID, count)
}

// Metrics_AddProviderCountForMarket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProviderCountForMarket'
type Metrics_AddProviderCountForMarket_Call struct {
	*mock.Call
}

// AddProviderCountForMarket is a helper method to define mock.On call
//   - pairID string
//   - count int
func (_e *Metrics_Expecter) AddProviderCountForMarket(pairID interface{}, count interface{}) *Metrics_AddProviderCountForMarket_Call {
	return &Metrics_AddProviderCountForMarket_Call{Call: _e.mock.On("AddProviderCountForMarket", pairID, count)}
}

func (_c *Metrics_AddProviderCountForMarket_Call) Run(run func(pairID string, count int)) *Metrics_AddProviderCountForMarket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *Metrics_AddProviderCountForMarket_Call) Return() *Metrics_AddProviderCountForMarket_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_AddProviderCountForMarket_Call) RunAndReturn(run func(string, int)) *Metrics_AddProviderCountForMarket_Call {
	_c.Call.Return(run)
	return _c
}

// AddProviderTick provides a mock function with given fields: providerName, pairID, success
func (_m *Metrics) AddProviderTick(providerName string, pairID string, success bool) {
	_m.Called(providerName, pairID, success)
}

// Metrics_AddProviderTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProviderTick'
type Metrics_AddProviderTick_Call struct {
	*mock.Call
}

// AddProviderTick is a helper method to define mock.On call
//   - providerName string
//   - pairID string
//   - success bool
func (_e *Metrics_Expecter) AddProviderTick(providerName interface{}, pairID interface{}, success interface{}) *Metrics_AddProviderTick_Call {
	return &Metrics_AddProviderTick_Call{Call: _e.mock.On("AddProviderTick", providerName, pairID, success)}
}

func (_c *Metrics_AddProviderTick_Call) Run(run func(providerName string, pairID string, success bool)) *Metrics_AddProviderTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *Metrics_AddProviderTick_Call) Return() *Metrics_AddProviderTick_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_AddProviderTick_Call) RunAndReturn(run func(string, string, bool)) *Metrics_AddProviderTick_Call {
	_c.Call.Return(run)
	return _c
}

// AddTick provides a mock function with given fields:
func (_m *Metrics) AddTick() {
	_m.Called()
}

// Metrics_AddTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTick'
type Metrics_AddTick_Call struct {
	*mock.Call
}

// AddTick is a helper method to define mock.On call
func (_e *Metrics_Expecter) AddTick() *Metrics_AddTick_Call {
	return &Metrics_AddTick_Call{Call: _e.mock.On("AddTick")}
}

func (_c *Metrics_AddTick_Call) Run(run func()) *Metrics_AddTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Metrics_AddTick_Call) Return() *Metrics_AddTick_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_AddTick_Call) RunAndReturn(run func()) *Metrics_AddTick_Call {
	_c.Call.Return(run)
	return _c
}

// AddTickerTick provides a mock function with given fields: pairID
func (_m *Metrics) AddTickerTick(pairID string) {
	_m.Called(pairID)
}

// Metrics_AddTickerTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTickerTick'
type Metrics_AddTickerTick_Call struct {
	*mock.Call
}

// AddTickerTick is a helper method to define mock.On call
//   - pairID string
func (_e *Metrics_Expecter) AddTickerTick(pairID interface{}) *Metrics_AddTickerTick_Call {
	return &Metrics_AddTickerTick_Call{Call: _e.mock.On("AddTickerTick", pairID)}
}

func (_c *Metrics_AddTickerTick_Call) Run(run func(pairID string)) *Metrics_AddTickerTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Metrics_AddTickerTick_Call) Return() *Metrics_AddTickerTick_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_AddTickerTick_Call) RunAndReturn(run func(string)) *Metrics_AddTickerTick_Call {
	_c.Call.Return(run)
	return _c
}

// GetMissingPrices provides a mock function with given fields:
func (_m *Metrics) GetMissingPrices() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMissingPrices")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Metrics_GetMissingPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMissingPrices'
type Metrics_GetMissingPrices_Call struct {
	*mock.Call
}

// GetMissingPrices is a helper method to define mock.On call
func (_e *Metrics_Expecter) GetMissingPrices() *Metrics_GetMissingPrices_Call {
	return &Metrics_GetMissingPrices_Call{Call: _e.mock.On("GetMissingPrices")}
}

func (_c *Metrics_GetMissingPrices_Call) Run(run func()) *Metrics_GetMissingPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Metrics_GetMissingPrices_Call) Return(_a0 []string) *Metrics_GetMissingPrices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Metrics_GetMissingPrices_Call) RunAndReturn(run func() []string) *Metrics_GetMissingPrices_Call {
	_c.Call.Return(run)
	return _c
}

// MissingPrices provides a mock function with given fields: pairIDs
func (_m *Metrics) MissingPrices(pairIDs []string) {
	_m.Called(pairIDs)
}

// Metrics_MissingPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MissingPrices'
type Metrics_MissingPrices_Call struct {
	*mock.Call
}

// MissingPrices is a helper method to define mock.On call
//   - pairIDs []string
func (_e *Metrics_Expecter) MissingPrices(pairIDs interface{}) *Metrics_MissingPrices_Call {
	return &Metrics_MissingPrices_Call{Call: _e.mock.On("MissingPrices", pairIDs)}
}

func (_c *Metrics_MissingPrices_Call) Run(run func(pairIDs []string)) *Metrics_MissingPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Metrics_MissingPrices_Call) Return() *Metrics_MissingPrices_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_MissingPrices_Call) RunAndReturn(run func([]string)) *Metrics_MissingPrices_Call {
	_c.Call.Return(run)
	return _c
}

// SetConnectBuildInfo provides a mock function with given fields:
func (_m *Metrics) SetConnectBuildInfo() {
	_m.Called()
}

// Metrics_SetConnectBuildInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConnectBuildInfo'
type Metrics_SetConnectBuildInfo_Call struct {
	*mock.Call
}

// SetConnectBuildInfo is a helper method to define mock.On call
func (_e *Metrics_Expecter) SetConnectBuildInfo() *Metrics_SetConnectBuildInfo_Call {
	return &Metrics_SetConnectBuildInfo_Call{Call: _e.mock.On("SetConnectBuildInfo")}
}

func (_c *Metrics_SetConnectBuildInfo_Call) Run(run func()) *Metrics_SetConnectBuildInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Metrics_SetConnectBuildInfo_Call) Return() *Metrics_SetConnectBuildInfo_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_SetConnectBuildInfo_Call) RunAndReturn(run func()) *Metrics_SetConnectBuildInfo_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAggregatePrice provides a mock function with given fields: pairID, decimals, price
func (_m *Metrics) UpdateAggregatePrice(pairID string, decimals uint64, price float64) {
	_m.Called(pairID, decimals, price)
}

// Metrics_UpdateAggregatePrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAggregatePrice'
type Metrics_UpdateAggregatePrice_Call struct {
	*mock.Call
}

// UpdateAggregatePrice is a helper method to define mock.On call
//   - pairID string
//   - decimals uint64
//   - price float64
func (_e *Metrics_Expecter) UpdateAggregatePrice(pairID interface{}, decimals interface{}, price interface{}) *Metrics_UpdateAggregatePrice_Call {
	return &Metrics_UpdateAggregatePrice_Call{Call: _e.mock.On("UpdateAggregatePrice", pairID, decimals, price)}
}

func (_c *Metrics_UpdateAggregatePrice_Call) Run(run func(pairID string, decimals uint64, price float64)) *Metrics_UpdateAggregatePrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64), args[2].(float64))
	})
	return _c
}

func (_c *Metrics_UpdateAggregatePrice_Call) Return() *Metrics_UpdateAggregatePrice_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_UpdateAggregatePrice_Call) RunAndReturn(run func(string, uint64, float64)) *Metrics_UpdateAggregatePrice_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePrice provides a mock function with given fields: name, pairID, decimals, price
func (_m *Metrics) UpdatePrice(name string, pairID string, decimals uint64, price float64) {
	_m.Called(name, pairID, decimals, price)
}

// Metrics_UpdatePrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePrice'
type Metrics_UpdatePrice_Call struct {
	*mock.Call
}

// UpdatePrice is a helper method to define mock.On call
//   - name string
//   - pairID string
//   - decimals uint64
//   - price float64
func (_e *Metrics_Expecter) UpdatePrice(name interface{}, pairID interface{}, decimals interface{}, price interface{}) *Metrics_UpdatePrice_Call {
	return &Metrics_UpdatePrice_Call{Call: _e.mock.On("UpdatePrice", name, pairID, decimals, price)}
}

func (_c *Metrics_UpdatePrice_Call) Run(run func(name string, pairID string, decimals uint64, price float64)) *Metrics_UpdatePrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(uint64), args[3].(float64))
	})
	return _c
}

func (_c *Metrics_UpdatePrice_Call) Return() *Metrics_UpdatePrice_Call {
	_c.Call.Return()
	return _c
}

func (_c *Metrics_UpdatePrice_Call) RunAndReturn(run func(string, string, uint64, float64)) *Metrics_UpdatePrice_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetrics creates a new instance of Metrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *Metrics {
	mock := &Metrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
