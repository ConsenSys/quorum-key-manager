// Code generated by MockGen. DO NOT EDIT.
// Source: vaults.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	"github.com/consensys/quorum-key-manager/src/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVaults is a mock of Vaults interface
type MockVaults struct {
	ctrl     *gomock.Controller
	recorder *MockVaultsMockRecorder
}

// MockVaultsMockRecorder is the mock recorder for MockVaults
type MockVaultsMockRecorder struct {
	mock *MockVaults
}

// NewMockVaults creates a new mock instance
func NewMockVaults(ctrl *gomock.Controller) *MockVaults {
	mock := &MockVaults{ctrl: ctrl}
	mock.recorder = &MockVaultsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVaults) EXPECT() *MockVaultsMockRecorder {
	return m.recorder
}

// CreateHashicorp mocks base method
func (m *MockVaults) CreateHashicorp(ctx context.Context, name string, config *entities.HashicorpConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHashicorp", ctx, name, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHashicorp indicates an expected call of CreateHashicorp
func (mr *MockVaultsMockRecorder) CreateHashicorp(ctx, name, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHashicorp", reflect.TypeOf((*MockVaults)(nil).CreateHashicorp), ctx, name, config)
}

// CreateAzure mocks base method
func (m *MockVaults) CreateAzure(ctx context.Context, name string, config *entities.AzureConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAzure", ctx, name, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAzure indicates an expected call of CreateAzure
func (mr *MockVaultsMockRecorder) CreateAzure(ctx, name, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAzure", reflect.TypeOf((*MockVaults)(nil).CreateAzure), ctx, name, config)
}

// CreateAWS mocks base method
func (m *MockVaults) CreateAWS(ctx context.Context, name string, config *entities.AWSConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAWS", ctx, name, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAWS indicates an expected call of CreateAWS
func (mr *MockVaultsMockRecorder) CreateAWS(ctx, name, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAWS", reflect.TypeOf((*MockVaults)(nil).CreateAWS), ctx, name, config)
}

// Get mocks base method
func (m *MockVaults) Get(ctx context.Context, name string) (*entities.Vault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name)
	ret0, _ := ret[0].(*entities.Vault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVaultsMockRecorder) Get(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVaults)(nil).Get), ctx, name)
}