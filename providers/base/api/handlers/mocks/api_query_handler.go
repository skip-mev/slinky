// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/providers/types"
)

// APIQueryHandler is an autogenerated mock type for the APIQueryHandler type
type APIQueryHandler[K types.ResponseKey, V types.ResponseValue] struct {
	mock.Mock
}

type APIQueryHandler_Expecter[K types.ResponseKey, V types.ResponseValue] struct {
	mock *mock.Mock
}

func (_m *APIQueryHandler[K, V]) EXPECT() *APIQueryHandler_Expecter[K, V] {
	return &APIQueryHandler_Expecter[K, V]{mock: &_m.Mock}
}

// Query provides a mock function with given fields: ctx, ids, responseCh
func (_m *APIQueryHandler[K, V]) Query(ctx context.Context, ids []K, responseCh chan<- types.GetResponse[K, V]) {
	_m.Called(ctx, ids, responseCh)
}

// APIQueryHandler_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type APIQueryHandler_Query_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []K
//   - responseCh chan<- types.GetResponse[K,V]
func (_e *APIQueryHandler_Expecter[K, V]) Query(ctx interface{}, ids interface{}, responseCh interface{}) *APIQueryHandler_Query_Call[K, V] {
	return &APIQueryHandler_Query_Call[K, V]{Call: _e.mock.On("Query", ctx, ids, responseCh)}
}

func (_c *APIQueryHandler_Query_Call[K, V]) Run(run func(ctx context.Context, ids []K, responseCh chan<- types.GetResponse[K, V])) *APIQueryHandler_Query_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]K), args[2].(chan<- types.GetResponse[K, V]))
	})
	return _c
}

func (_c *APIQueryHandler_Query_Call[K, V]) Return() *APIQueryHandler_Query_Call[K, V] {
	_c.Call.Return()
	return _c
}

func (_c *APIQueryHandler_Query_Call[K, V]) RunAndReturn(run func(context.Context, []K, chan<- types.GetResponse[K, V])) *APIQueryHandler_Query_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// NewAPIQueryHandler creates a new instance of APIQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIQueryHandler[K types.ResponseKey, V types.ResponseValue](t interface {
	mock.TestingT
	Cleanup(func())
}) *APIQueryHandler[K, V] {
	mock := &APIQueryHandler[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
