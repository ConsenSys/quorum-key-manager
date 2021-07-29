// Code generated by MockGen. DO NOT EDIT.
// Source: eth1.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	ethereum "github.com/consensys/quorum-key-manager/pkg/ethereum"
	entities "github.com/consensys/quorum-key-manager/src/stores/store/entities"
	types "github.com/consensys/quorum/core/types"
	types0 "github.com/ethereum/go-ethereum/core/types"
	core "github.com/ethereum/go-ethereum/signer/core"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockStore) Info(arg0 context.Context) (*entities.StoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(*entities.StoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockStoreMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockStore)(nil).Info), arg0)
}

// Create mocks base method
func (m *MockStore) Create(ctx context.Context, id string, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, id, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(ctx, id, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), ctx, id, attr)
}

// Import mocks base method
func (m *MockStore) Import(ctx context.Context, id string, privKey []byte, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, id, privKey, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import
func (mr *MockStoreMockRecorder) Import(ctx, id, privKey, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockStore)(nil).Import), ctx, id, privKey, attr)
}

// Get mocks base method
func (m *MockStore) Get(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, addr)
}

// GetAll mocks base method
func (m *MockStore) GetAll(ctx context.Context) ([]*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockStoreMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStore)(nil).GetAll), ctx)
}

// List mocks base method
func (m *MockStore) List(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), ctx)
}

// Update mocks base method
func (m *MockStore) Update(ctx context.Context, addr string, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, addr, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(ctx, addr, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), ctx, addr, attr)
}

// Delete mocks base method
func (m *MockStore) Delete(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), ctx, addr)
}

// GetDeleted mocks base method
func (m *MockStore) GetDeleted(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleted", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeleted indicates an expected call of GetDeleted
func (mr *MockStoreMockRecorder) GetDeleted(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleted", reflect.TypeOf((*MockStore)(nil).GetDeleted), ctx, addr)
}

// ListDeleted mocks base method
func (m *MockStore) ListDeleted(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeleted", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeleted indicates an expected call of ListDeleted
func (mr *MockStoreMockRecorder) ListDeleted(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeleted", reflect.TypeOf((*MockStore)(nil).ListDeleted), ctx)
}

// Restore mocks base method
func (m *MockStore) Undelete(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockStoreMockRecorder) Undelete(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockStore)(nil).Undelete), ctx, addr)
}

// Destroy mocks base method
func (m *MockStore) Destroy(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockStoreMockRecorder) Destroy(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockStore)(nil).Destroy), ctx, addr)
}

// Sign mocks base method
func (m *MockStore) Sign(ctx context.Context, addr string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockStoreMockRecorder) Sign(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockStore)(nil).Sign), ctx, addr, data)
}

// SignData mocks base method
func (m *MockStore) SignData(ctx context.Context, addr string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignData", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignData indicates an expected call of SignData
func (mr *MockStoreMockRecorder) SignData(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignData", reflect.TypeOf((*MockStore)(nil).SignData), ctx, addr, data)
}

// SignTypedData mocks base method
func (m *MockStore) SignTypedData(ctx context.Context, addr string, typedData *core.TypedData) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTypedData", ctx, addr, typedData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTypedData indicates an expected call of SignTypedData
func (mr *MockStoreMockRecorder) SignTypedData(ctx, addr, typedData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTypedData", reflect.TypeOf((*MockStore)(nil).SignTypedData), ctx, addr, typedData)
}

// SignTransaction mocks base method
func (m *MockStore) SignTransaction(ctx context.Context, addr string, chainID *big.Int, tx *types0.Transaction) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction", ctx, addr, chainID, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction
func (mr *MockStoreMockRecorder) SignTransaction(ctx, addr, chainID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockStore)(nil).SignTransaction), ctx, addr, chainID, tx)
}

// SignEEA mocks base method
func (m *MockStore) SignEEA(ctx context.Context, addr string, chainID *big.Int, tx *types0.Transaction, args *ethereum.PrivateArgs) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignEEA", ctx, addr, chainID, tx, args)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignEEA indicates an expected call of SignEEA
func (mr *MockStoreMockRecorder) SignEEA(ctx, addr, chainID, tx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignEEA", reflect.TypeOf((*MockStore)(nil).SignEEA), ctx, addr, chainID, tx, args)
}

// SignPrivate mocks base method
func (m *MockStore) SignPrivate(ctx context.Context, addr string, tx *types.Transaction) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignPrivate", ctx, addr, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignPrivate indicates an expected call of SignPrivate
func (mr *MockStoreMockRecorder) SignPrivate(ctx, addr, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignPrivate", reflect.TypeOf((*MockStore)(nil).SignPrivate), ctx, addr, tx)
}

// ECRevocer mocks base method
func (m *MockStore) ECRevocer(ctx context.Context, data, sig []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ECRevocer", ctx, data, sig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ECRevocer indicates an expected call of ECRevocer
func (mr *MockStoreMockRecorder) ECRevocer(ctx, data, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ECRevocer", reflect.TypeOf((*MockStore)(nil).ECRevocer), ctx, data, sig)
}

// Verify mocks base method
func (m *MockStore) Verify(ctx context.Context, addr string, data, sig []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, addr, data, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockStoreMockRecorder) Verify(ctx, addr, data, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockStore)(nil).Verify), ctx, addr, data, sig)
}

// VerifyTypedData mocks base method
func (m *MockStore) VerifyTypedData(ctx context.Context, addr string, typedData *core.TypedData, sig []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTypedData", ctx, addr, typedData, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTypedData indicates an expected call of VerifyTypedData
func (mr *MockStoreMockRecorder) VerifyTypedData(ctx, addr, typedData, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTypedData", reflect.TypeOf((*MockStore)(nil).VerifyTypedData), ctx, addr, typedData, sig)
}

// Encrypt mocks base method
func (m *MockStore) Encrypt(ctx context.Context, addr string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockStoreMockRecorder) Encrypt(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockStore)(nil).Encrypt), ctx, addr, data)
}

// Decrypt mocks base method
func (m *MockStore) Decrypt(ctx context.Context, addr string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt
func (mr *MockStoreMockRecorder) Decrypt(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockStore)(nil).Decrypt), ctx, addr, data)
}
