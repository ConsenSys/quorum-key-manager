// Code generated by MockGen. DO NOT EDIT.
// Source: preparer.go

// Package request is a generated GoMock package.
package request

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPreparer is a mock of Preparer interface.
type MockPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockPreparerMockRecorder
}

// MockPreparerMockRecorder is the mock recorder for MockPreparer.
type MockPreparerMockRecorder struct {
	mock *MockPreparer
}

// NewMockPreparer creates a new mock instance.
func NewMockPreparer(ctrl *gomock.Controller) *MockPreparer {
	mock := &MockPreparer{ctrl: ctrl}
	mock.recorder = &MockPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreparer) EXPECT() *MockPreparerMockRecorder {
	return m.recorder
}

// Prepare mock base method.
func (m *MockPreparer) Prepare(arg0 *http.Request) (*http.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockPreparerMockRecorder) Prepare(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockPreparer)(nil).Prepare), arg0)
}
