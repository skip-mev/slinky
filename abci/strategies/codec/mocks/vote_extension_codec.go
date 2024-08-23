// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/abci/ve/types"
)

// VoteExtensionCodec is an autogenerated mock type for the VoteExtensionCodec type
type VoteExtensionCodec struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0
func (_m *VoteExtensionCodec) Decode(_a0 []byte) (types.OracleVoteExtension, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 types.OracleVoteExtension
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (types.OracleVoteExtension, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) types.OracleVoteExtension); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(types.OracleVoteExtension)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encode provides a mock function with given fields: ve
func (_m *VoteExtensionCodec) Encode(ve types.OracleVoteExtension) ([]byte, error) {
	ret := _m.Called(ve)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(types.OracleVoteExtension) ([]byte, error)); ok {
		return rf(ve)
	}
	if rf, ok := ret.Get(0).(func(types.OracleVoteExtension) []byte); ok {
		r0 = rf(ve)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(types.OracleVoteExtension) error); ok {
		r1 = rf(ve)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVoteExtensionCodec creates a new instance of VoteExtensionCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVoteExtensionCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *VoteExtensionCodec {
	mock := &VoteExtensionCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
