// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/x/marketmap/types"
)

// QueryClient is an autogenerated mock type for the QueryClient type
type QueryClient struct {
	mock.Mock
}

// LastUpdated provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) LastUpdated(ctx context.Context, in *types.LastUpdatedRequest, opts ...grpc.CallOption) (*types.LastUpdatedResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.LastUpdatedResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.LastUpdatedRequest, ...grpc.CallOption) (*types.LastUpdatedResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.LastUpdatedRequest, ...grpc.CallOption) *types.LastUpdatedResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.LastUpdatedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.LastUpdatedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Market provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Market(ctx context.Context, in *types.MarketRequest, opts ...grpc.CallOption) (*types.MarketResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.MarketResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.MarketRequest, ...grpc.CallOption) (*types.MarketResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.MarketRequest, ...grpc.CallOption) *types.MarketResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MarketResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.MarketRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketMap provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) MarketMap(ctx context.Context, in *types.MarketMapRequest, opts ...grpc.CallOption) (*types.MarketMapResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.MarketMapResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.MarketMapRequest, ...grpc.CallOption) (*types.MarketMapResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.MarketMapRequest, ...grpc.CallOption) *types.MarketMapResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MarketMapResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.MarketMapRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *QueryClient) Params(ctx context.Context, in *types.ParamsRequest, opts ...grpc.CallOption) (*types.ParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.ParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.ParamsRequest, ...grpc.CallOption) (*types.ParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.ParamsRequest, ...grpc.CallOption) *types.ParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.ParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryClient creates a new instance of QueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryClient {
	mock := &QueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
