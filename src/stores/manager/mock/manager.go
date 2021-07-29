// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/consensys/quorum-key-manager/src/auth/types"
	manifest "github.com/consensys/quorum-key-manager/src/manifests/types"
	connectors "github.com/consensys/quorum-key-manager/src/stores/connectors"
	entities "github.com/consensys/quorum-key-manager/src/stores/store/entities"
	eth1 "github.com/consensys/quorum-key-manager/src/stores/store/eth1"
	secrets "github.com/consensys/quorum-key-manager/src/stores/store/secrets"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetSecretStore mocks base method
func (m *MockManager) GetSecretStore(ctx context.Context, name string, userInfo *types.UserInfo) (secrets.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretStore", ctx, name, userInfo)
	ret0, _ := ret[0].(secrets.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretStore indicates an expected call of GetSecretStore
func (mr *MockManagerMockRecorder) GetSecretStore(ctx, name, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretStore", reflect.TypeOf((*MockManager)(nil).GetSecretStore), ctx, name, userInfo)
}

// GetKeyStore mocks base method
func (m *MockManager) GetKeyStore(ctx context.Context, name string, userInfo *types.UserInfo) (connectors.KeysConnector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyStore", ctx, name, userInfo)
	ret0, _ := ret[0].(connectors.KeysConnector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyStore indicates an expected call of GetKeyStore
func (mr *MockManagerMockRecorder) GetKeyStore(ctx, name, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyStore", reflect.TypeOf((*MockManager)(nil).GetKeyStore), ctx, name, userInfo)
}

// GetEth1Store mocks base method
func (m *MockManager) GetEth1Store(ctx context.Context, name string, userInfo *types.UserInfo) (eth1.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEth1Store", ctx, name, userInfo)
	ret0, _ := ret[0].(eth1.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEth1Store indicates an expected call of GetEth1Store
func (mr *MockManagerMockRecorder) GetEth1Store(ctx, name, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEth1Store", reflect.TypeOf((*MockManager)(nil).GetEth1Store), ctx, name, userInfo)
}

// GetEth1StoreByAddr mocks base method
func (m *MockManager) GetEth1StoreByAddr(ctx context.Context, addr common.Address, userInfo *types.UserInfo) (eth1.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEth1StoreByAddr", ctx, addr, userInfo)
	ret0, _ := ret[0].(eth1.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEth1StoreByAddr indicates an expected call of GetEth1StoreByAddr
func (mr *MockManagerMockRecorder) GetEth1StoreByAddr(ctx, addr, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEth1StoreByAddr", reflect.TypeOf((*MockManager)(nil).GetEth1StoreByAddr), ctx, addr, userInfo)
}

// List mocks base method
func (m *MockManager) List(ctx context.Context, kind manifest.Kind, userInfo *types.UserInfo) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, kind, userInfo)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockManagerMockRecorder) List(ctx, kind, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), ctx, kind, userInfo)
}

// ListAllAccounts mocks base method
func (m *MockManager) ListAllAccounts(ctx context.Context, userInfo *types.UserInfo) ([]*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllAccounts", ctx, userInfo)
	ret0, _ := ret[0].([]*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllAccounts indicates an expected call of ListAllAccounts
func (mr *MockManagerMockRecorder) ListAllAccounts(ctx, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllAccounts", reflect.TypeOf((*MockManager)(nil).ListAllAccounts), ctx, userInfo)
}
