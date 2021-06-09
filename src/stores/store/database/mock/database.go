// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	database "github.com/ConsenSysQuorum/quorum-key-manager/src/stores/store/database"
	entities "github.com/ConsenSysQuorum/quorum-key-manager/src/stores/store/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// ETH1Accounts mocks base method
func (m *MockDatabase) ETH1Accounts() database.ETH1Accounts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ETH1Accounts")
	ret0, _ := ret[0].(database.ETH1Accounts)
	return ret0
}

// ETH1Accounts indicates an expected call of ETH1Accounts
func (mr *MockDatabaseMockRecorder) ETH1Accounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ETH1Accounts", reflect.TypeOf((*MockDatabase)(nil).ETH1Accounts))
}

// MockETH1Accounts is a mock of ETH1Accounts interface
type MockETH1Accounts struct {
	ctrl     *gomock.Controller
	recorder *MockETH1AccountsMockRecorder
}

// MockETH1AccountsMockRecorder is the mock recorder for MockETH1Accounts
type MockETH1AccountsMockRecorder struct {
	mock *MockETH1Accounts
}

// NewMockETH1Accounts creates a new mock instance
func NewMockETH1Accounts(ctrl *gomock.Controller) *MockETH1Accounts {
	mock := &MockETH1Accounts{ctrl: ctrl}
	mock.recorder = &MockETH1AccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockETH1Accounts) EXPECT() *MockETH1AccountsMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockETH1Accounts) Get(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockETH1AccountsMockRecorder) Get(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockETH1Accounts)(nil).Get), ctx, addr)
}

// GetDeleted mocks base method
func (m *MockETH1Accounts) GetDeleted(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleted", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeleted indicates an expected call of GetDeleted
func (mr *MockETH1AccountsMockRecorder) GetDeleted(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleted", reflect.TypeOf((*MockETH1Accounts)(nil).GetDeleted), ctx, addr)
}

// GetAll mocks base method
func (m *MockETH1Accounts) GetAll(ctx context.Context) ([]*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockETH1AccountsMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockETH1Accounts)(nil).GetAll), ctx)
}

// GetAllDeleted mocks base method
func (m *MockETH1Accounts) GetAllDeleted(ctx context.Context) ([]*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDeleted", ctx)
	ret0, _ := ret[0].([]*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDeleted indicates an expected call of GetAllDeleted
func (mr *MockETH1AccountsMockRecorder) GetAllDeleted(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDeleted", reflect.TypeOf((*MockETH1Accounts)(nil).GetAllDeleted), ctx)
}

// Add mocks base method
func (m *MockETH1Accounts) Add(ctx context.Context, account *entities.ETH1Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockETH1AccountsMockRecorder) Add(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockETH1Accounts)(nil).Add), ctx, account)
}

// AddDeleted mocks base method
func (m *MockETH1Accounts) AddDeleted(ctx context.Context, account *entities.ETH1Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeleted", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDeleted indicates an expected call of AddDeleted
func (mr *MockETH1AccountsMockRecorder) AddDeleted(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeleted", reflect.TypeOf((*MockETH1Accounts)(nil).AddDeleted), ctx, account)
}

// Remove mocks base method
func (m *MockETH1Accounts) Remove(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockETH1AccountsMockRecorder) Remove(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockETH1Accounts)(nil).Remove), ctx, addr)
}

// RemoveDeleted mocks base method
func (m *MockETH1Accounts) RemoveDeleted(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDeleted", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDeleted indicates an expected call of RemoveDeleted
func (mr *MockETH1AccountsMockRecorder) RemoveDeleted(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDeleted", reflect.TypeOf((*MockETH1Accounts)(nil).RemoveDeleted), ctx, addr)
}
