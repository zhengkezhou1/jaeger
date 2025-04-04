// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uimodel "github.com/jaegertracing/jaeger/internal/uimodel"
)

// QueryService is an autogenerated mock type for the QueryService type
type QueryService struct {
	mock.Mock
}

// GetTraces provides a mock function with given fields: serviceName, operation, tags
func (_m *QueryService) GetTraces(serviceName string, operation string, tags map[string]string) ([]*uimodel.Trace, error) {
	ret := _m.Called(serviceName, operation, tags)

	if len(ret) == 0 {
		panic("no return value specified for GetTraces")
	}

	var r0 []*uimodel.Trace
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) ([]*uimodel.Trace, error)); ok {
		return rf(serviceName, operation, tags)
	}
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) []*uimodel.Trace); ok {
		r0 = rf(serviceName, operation, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*uimodel.Trace)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, map[string]string) error); ok {
		r1 = rf(serviceName, operation, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryService creates a new instance of QueryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryService {
	mock := &QueryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
