// Code generated by MockGen. DO NOT EDIT.
// Source: internal/wrap/manet.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// MockManeter is a mock of Maneter interface.
type MockManeter struct {
	ctrl     *gomock.Controller
	recorder *MockManeterMockRecorder
}

// MockManeterMockRecorder is the mock recorder for MockManeter.
type MockManeterMockRecorder struct {
	mock *MockManeter
}

// NewMockManeter creates a new mock instance.
func NewMockManeter(ctrl *gomock.Controller) *MockManeter {
	mock := &MockManeter{ctrl: ctrl}
	mock.recorder = &MockManeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManeter) EXPECT() *MockManeterMockRecorder {
	return m.recorder
}

// IsPublicAddr mocks base method.
func (m *MockManeter) IsPublicAddr(a multiaddr.Multiaddr) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPublicAddr", a)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPublicAddr indicates an expected call of IsPublicAddr.
func (mr *MockManeterMockRecorder) IsPublicAddr(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPublicAddr", reflect.TypeOf((*MockManeter)(nil).IsPublicAddr), a)
}
