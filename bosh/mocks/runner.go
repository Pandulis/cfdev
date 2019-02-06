// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/bosh (interfaces: Runner)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	exec "os/exec"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Output mocks base method
func (m *MockRunner) Output(arg0 *exec.Cmd) ([]byte, error) {
	ret := m.ctrl.Call(m, "Output", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output
func (mr *MockRunnerMockRecorder) Output(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockRunner)(nil).Output), arg0)
}