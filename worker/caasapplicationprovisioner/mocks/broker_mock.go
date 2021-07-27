// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasapplicationprovisioner (interfaces: CAASBroker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	caas "github.com/juju/juju/caas"
	names "github.com/juju/names/v4"
)

// MockCAASBroker is a mock of CAASBroker interface.
type MockCAASBroker struct {
	ctrl     *gomock.Controller
	recorder *MockCAASBrokerMockRecorder
}

// MockCAASBrokerMockRecorder is the mock recorder for MockCAASBroker.
type MockCAASBrokerMockRecorder struct {
	mock *MockCAASBroker
}

// NewMockCAASBroker creates a new mock instance.
func NewMockCAASBroker(ctrl *gomock.Controller) *MockCAASBroker {
	mock := &MockCAASBroker{ctrl: ctrl}
	mock.recorder = &MockCAASBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASBroker) EXPECT() *MockCAASBrokerMockRecorder {
	return m.recorder
}

// AnnotateUnit mocks base method.
func (m *MockCAASBroker) AnnotateUnit(arg0 string, arg1 caas.DeploymentMode, arg2 string, arg3 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotateUnit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnotateUnit indicates an expected call of AnnotateUnit.
func (mr *MockCAASBrokerMockRecorder) AnnotateUnit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotateUnit", reflect.TypeOf((*MockCAASBroker)(nil).AnnotateUnit), arg0, arg1, arg2, arg3)
}

// Application mocks base method.
func (m *MockCAASBroker) Application(arg0 string, arg1 caas.DeploymentType) caas.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0, arg1)
	ret0, _ := ret[0].(caas.Application)
	return ret0
}

// Application indicates an expected call of Application.
func (mr *MockCAASBrokerMockRecorder) Application(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockCAASBroker)(nil).Application), arg0, arg1)
}
