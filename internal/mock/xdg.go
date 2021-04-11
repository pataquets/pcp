// Code generated by MockGen. DO NOT EDIT.
// Source: internal/wrap/xdg.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockXdger is a mock of Xdger interface.
type MockXdger struct {
	ctrl     *gomock.Controller
	recorder *MockXdgerMockRecorder
}

// MockXdgerMockRecorder is the mock recorder for MockXdger.
type MockXdgerMockRecorder struct {
	mock *MockXdger
}

// NewMockXdger creates a new mock instance.
func NewMockXdger(ctrl *gomock.Controller) *MockXdger {
	mock := &MockXdger{ctrl: ctrl}
	mock.recorder = &MockXdgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXdger) EXPECT() *MockXdgerMockRecorder {
	return m.recorder
}

// ConfigFile mocks base method.
func (m *MockXdger) ConfigFile(relPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFile", relPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigFile indicates an expected call of ConfigFile.
func (mr *MockXdgerMockRecorder) ConfigFile(relPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFile", reflect.TypeOf((*MockXdger)(nil).ConfigFile), relPath)
}
