// Code generated by MockGen. DO NOT EDIT.
// Source: ./example.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExample is a mock of Example interface
type MockExample struct {
	ctrl     *gomock.Controller
	recorder *MockExampleMockRecorder
}

// MockExampleMockRecorder is the mock recorder for MockExample
type MockExampleMockRecorder struct {
	mock *MockExample
}

// NewMockExample creates a new mock instance
func NewMockExample(ctrl *gomock.Controller) *MockExample {
	mock := &MockExample{ctrl: ctrl}
	mock.recorder = &MockExampleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExample) EXPECT() *MockExampleMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockExample) Get(key string) (string, error) {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockExampleMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockExample)(nil).Get), key)
}

// Set mocks base method
func (m *MockExample) Set(key, value string) error {
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockExampleMockRecorder) Set(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockExample)(nil).Set), key, value)
}

// Test mocks base method
func (m *MockExample) Test() {
	m.ctrl.Call(m, "Test")
}

// Test indicates an expected call of Test
func (mr *MockExampleMockRecorder) Test() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockExample)(nil).Test))
}
