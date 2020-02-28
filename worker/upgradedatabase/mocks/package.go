// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/upgradedatabase (interfaces: Logger,Pool,UpgradeInfo)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	upgradedatabase "github.com/juju/juju/worker/upgradedatabase"
	version "github.com/juju/version"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Infof mocks base method
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// MockPool is a mock of Pool interface
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPool) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPoolMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPool)(nil).Close))
}

// EnsureUpgradeInfo mocks base method
func (m *MockPool) EnsureUpgradeInfo(arg0 string, arg1, arg2 version.Number) (upgradedatabase.UpgradeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureUpgradeInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(upgradedatabase.UpgradeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureUpgradeInfo indicates an expected call of EnsureUpgradeInfo
func (mr *MockPoolMockRecorder) EnsureUpgradeInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureUpgradeInfo", reflect.TypeOf((*MockPool)(nil).EnsureUpgradeInfo), arg0, arg1, arg2)
}

// IsPrimary mocks base method
func (m *MockPool) IsPrimary(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrimary", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPrimary indicates an expected call of IsPrimary
func (mr *MockPoolMockRecorder) IsPrimary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrimary", reflect.TypeOf((*MockPool)(nil).IsPrimary), arg0)
}

// SetStatus mocks base method
func (m *MockPool) SetStatus(arg0 string, arg1 status.Status, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockPoolMockRecorder) SetStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockPool)(nil).SetStatus), arg0, arg1, arg2)
}

// MockUpgradeInfo is a mock of UpgradeInfo interface
type MockUpgradeInfo struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeInfoMockRecorder
}

// MockUpgradeInfoMockRecorder is the mock recorder for MockUpgradeInfo
type MockUpgradeInfoMockRecorder struct {
	mock *MockUpgradeInfo
}

// NewMockUpgradeInfo creates a new mock instance
func NewMockUpgradeInfo(ctrl *gomock.Controller) *MockUpgradeInfo {
	mock := &MockUpgradeInfo{ctrl: ctrl}
	mock.recorder = &MockUpgradeInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpgradeInfo) EXPECT() *MockUpgradeInfoMockRecorder {
	return m.recorder
}

// Refresh mocks base method
func (m *MockUpgradeInfo) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockUpgradeInfoMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockUpgradeInfo)(nil).Refresh))
}

// SetStatus mocks base method
func (m *MockUpgradeInfo) SetStatus(arg0 state.UpgradeStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockUpgradeInfoMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockUpgradeInfo)(nil).SetStatus), arg0)
}

// Status mocks base method
func (m *MockUpgradeInfo) Status() state.UpgradeStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(state.UpgradeStatus)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockUpgradeInfoMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockUpgradeInfo)(nil).Status))
}

// Watch mocks base method
func (m *MockUpgradeInfo) Watch() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockUpgradeInfoMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockUpgradeInfo)(nil).Watch))
}
