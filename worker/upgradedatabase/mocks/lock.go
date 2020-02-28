// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/gate (interfaces: Lock)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLock is a mock of Lock interface
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// IsUnlocked mocks base method
func (m *MockLock) IsUnlocked() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnlocked")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnlocked indicates an expected call of IsUnlocked
func (mr *MockLockMockRecorder) IsUnlocked() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnlocked", reflect.TypeOf((*MockLock)(nil).IsUnlocked))
}

// Unlock mocks base method
func (m *MockLock) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock
func (mr *MockLockMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockLock)(nil).Unlock))
}

// Unlocked mocks base method
func (m *MockLock) Unlocked() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlocked")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Unlocked indicates an expected call of Unlocked
func (mr *MockLockMockRecorder) Unlocked() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlocked", reflect.TypeOf((*MockLock)(nil).Unlocked))
}
