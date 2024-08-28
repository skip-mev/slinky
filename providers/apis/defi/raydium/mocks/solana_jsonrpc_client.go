// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/gagliardetto/solana-go/rpc"

	solana "github.com/gagliardetto/solana-go"
)

// SolanaJSONRPCClient is an autogenerated mock type for the SolanaJSONRPCClient type
type SolanaJSONRPCClient struct {
	mock.Mock
}

type SolanaJSONRPCClient_Expecter struct {
	mock *mock.Mock
}

func (_m *SolanaJSONRPCClient) EXPECT() *SolanaJSONRPCClient_Expecter {
	return &SolanaJSONRPCClient_Expecter{mock: &_m.Mock}
}

// GetMultipleAccountsWithOpts provides a mock function with given fields: ctx, accounts, opts
func (_m *SolanaJSONRPCClient) GetMultipleAccountsWithOpts(ctx context.Context, accounts []solana.PublicKey, opts *rpc.GetMultipleAccountsOpts) (*rpc.GetMultipleAccountsResult, error) {
	ret := _m.Called(ctx, accounts, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetMultipleAccountsWithOpts")
	}

	var r0 *rpc.GetMultipleAccountsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []solana.PublicKey, *rpc.GetMultipleAccountsOpts) (*rpc.GetMultipleAccountsResult, error)); ok {
		return rf(ctx, accounts, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []solana.PublicKey, *rpc.GetMultipleAccountsOpts) *rpc.GetMultipleAccountsResult); ok {
		r0 = rf(ctx, accounts, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetMultipleAccountsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []solana.PublicKey, *rpc.GetMultipleAccountsOpts) error); ok {
		r1 = rf(ctx, accounts, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMultipleAccountsWithOpts'
type SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call struct {
	*mock.Call
}

// GetMultipleAccountsWithOpts is a helper method to define mock.On call
//   - ctx context.Context
//   - accounts []solana.PublicKey
//   - opts *rpc.GetMultipleAccountsOpts
func (_e *SolanaJSONRPCClient_Expecter) GetMultipleAccountsWithOpts(ctx interface{}, accounts interface{}, opts interface{}) *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call {
	return &SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call{Call: _e.mock.On("GetMultipleAccountsWithOpts", ctx, accounts, opts)}
}

func (_c *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call) Run(run func(ctx context.Context, accounts []solana.PublicKey, opts *rpc.GetMultipleAccountsOpts)) *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]solana.PublicKey), args[2].(*rpc.GetMultipleAccountsOpts))
	})
	return _c
}

func (_c *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call) Return(out *rpc.GetMultipleAccountsResult, err error) *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call) RunAndReturn(run func(context.Context, []solana.PublicKey, *rpc.GetMultipleAccountsOpts) (*rpc.GetMultipleAccountsResult, error)) *SolanaJSONRPCClient_GetMultipleAccountsWithOpts_Call {
	_c.Call.Return(run)
	return _c
}

// NewSolanaJSONRPCClient creates a new instance of SolanaJSONRPCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSolanaJSONRPCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SolanaJSONRPCClient {
	mock := &SolanaJSONRPCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
