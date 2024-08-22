// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ethereum/go-ethereum/rpc"
)

// EVMClient is an autogenerated mock type for the EVMClient type
type EVMClient struct {
	mock.Mock
}

// BatchCallContext provides a mock function with given fields: ctx, calls
func (_m *EVMClient) BatchCallContext(ctx context.Context, calls []rpc.BatchElem) error {
	ret := _m.Called(ctx, calls)

	if len(ret) == 0 {
		panic("no return value specified for BatchCallContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, calls)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEVMClient creates a new instance of EVMClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEVMClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EVMClient {
	mock := &EVMClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
