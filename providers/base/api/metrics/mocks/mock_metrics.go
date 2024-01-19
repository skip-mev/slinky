// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	metrics "github.com/skip-mev/slinky/providers/base/api/metrics"

	time "time"
)

// APIMetrics is an autogenerated mock type for the APIMetrics type
type APIMetrics struct {
	mock.Mock
}

// AddProviderResponse provides a mock function with given fields: providerName, id, status
func (_m *APIMetrics) AddProviderResponse(providerName string, id string, status metrics.Status) {
	_m.Called(providerName, id, status)
}

// ObserveProviderResponseLatency provides a mock function with given fields: providerName, duration
func (_m *APIMetrics) ObserveProviderResponseLatency(providerName string, duration time.Duration) {
	_m.Called(providerName, duration)
}

// NewAPIMetrics creates a new instance of APIMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIMetrics(t interface {
	mock.TestingT
	Cleanup(func())
},
) *APIMetrics {
	mock := &APIMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
