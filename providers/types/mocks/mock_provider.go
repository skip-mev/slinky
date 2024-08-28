// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/providers/types"
)

// Provider is an autogenerated mock type for the Provider type
type Provider[K types.ResponseKey, V types.ResponseValue] struct {
	mock.Mock
}

type Provider_Expecter[K types.ResponseKey, V types.ResponseValue] struct {
	mock *mock.Mock
}

func (_m *Provider[K, V]) EXPECT() *Provider_Expecter[K, V] {
	return &Provider_Expecter[K, V]{mock: &_m.Mock}
}

// GetData provides a mock function with given fields:
func (_m *Provider[K, V]) GetData() map[K]types.ResolvedResult[V] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 map[K]types.ResolvedResult[V]
	if rf, ok := ret.Get(0).(func() map[K]types.ResolvedResult[V]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[K]types.ResolvedResult[V])
		}
	}

	return r0
}

// Provider_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type Provider_GetData_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
func (_e *Provider_Expecter[K, V]) GetData() *Provider_GetData_Call[K, V] {
	return &Provider_GetData_Call[K, V]{Call: _e.mock.On("GetData")}
}

func (_c *Provider_GetData_Call[K, V]) Run(run func()) *Provider_GetData_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Provider_GetData_Call[K, V]) Return(_a0 map[K]types.ResolvedResult[V]) *Provider_GetData_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_GetData_Call[K, V]) RunAndReturn(run func() map[K]types.ResolvedResult[V]) *Provider_GetData_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *Provider[K, V]) IsRunning() bool {
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

// Provider_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type Provider_IsRunning_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *Provider_Expecter[K, V]) IsRunning() *Provider_IsRunning_Call[K, V] {
	return &Provider_IsRunning_Call[K, V]{Call: _e.mock.On("IsRunning")}
}

func (_c *Provider_IsRunning_Call[K, V]) Run(run func()) *Provider_IsRunning_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Provider_IsRunning_Call[K, V]) Return(_a0 bool) *Provider_IsRunning_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_IsRunning_Call[K, V]) RunAndReturn(run func() bool) *Provider_IsRunning_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Provider[K, V]) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Provider_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Provider_Name_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Provider_Expecter[K, V]) Name() *Provider_Name_Call[K, V] {
	return &Provider_Name_Call[K, V]{Call: _e.mock.On("Name")}
}

func (_c *Provider_Name_Call[K, V]) Run(run func()) *Provider_Name_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Provider_Name_Call[K, V]) Return(_a0 string) *Provider_Name_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_Name_Call[K, V]) RunAndReturn(run func() string) *Provider_Name_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Provider[K, V]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Provider_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Provider_Start_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Provider_Expecter[K, V]) Start(_a0 interface{}) *Provider_Start_Call[K, V] {
	return &Provider_Start_Call[K, V]{Call: _e.mock.On("Start", _a0)}
}

func (_c *Provider_Start_Call[K, V]) Run(run func(_a0 context.Context)) *Provider_Start_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Provider_Start_Call[K, V]) Return(_a0 error) *Provider_Start_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_Start_Call[K, V]) RunAndReturn(run func(context.Context) error) *Provider_Start_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *Provider[K, V]) Type() types.ProviderType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 types.ProviderType
	if rf, ok := ret.Get(0).(func() types.ProviderType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.ProviderType)
	}

	return r0
}

// Provider_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type Provider_Type_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *Provider_Expecter[K, V]) Type() *Provider_Type_Call[K, V] {
	return &Provider_Type_Call[K, V]{Call: _e.mock.On("Type")}
}

func (_c *Provider_Type_Call[K, V]) Run(run func()) *Provider_Type_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Provider_Type_Call[K, V]) Return(_a0 types.ProviderType) *Provider_Type_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_Type_Call[K, V]) RunAndReturn(run func() types.ProviderType) *Provider_Type_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvider[K types.ResponseKey, V types.ResponseValue](t interface {
	mock.TestingT
	Cleanup(func())
}) *Provider[K, V] {
	mock := &Provider[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
