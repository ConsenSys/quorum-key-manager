// Code generated by MockGen. DO NOT EDIT.
// Source: eth1.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	ethereum "github.com/consensys/quorum-key-manager/pkg/ethereum"
	entities "github.com/consensys/quorum-key-manager/src/stores/entities"
	types "github.com/consensys/quorum/core/types"
	common "github.com/ethereum/go-ethereum/common"
	types0 "github.com/ethereum/go-ethereum/core/types"
	core "github.com/ethereum/go-ethereum/signer/core"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockEth1Store is a mock of Eth1Store interface
type MockEth1Store struct {
	ctrl     *gomock.Controller
	recorder *MockEth1StoreMockRecorder
}

// MockEth1StoreMockRecorder is the mock recorder for MockEth1Store
type MockEth1StoreMockRecorder struct {
	mock *MockEth1Store
}

// NewMockEth1Store creates a new mock instance
func NewMockEth1Store(ctrl *gomock.Controller) *MockEth1Store {
	mock := &MockEth1Store{ctrl: ctrl}
	mock.recorder = &MockEth1StoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEth1Store) EXPECT() *MockEth1StoreMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockEth1Store) Info(arg0 context.Context) (*entities.StoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(*entities.StoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockEth1StoreMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockEth1Store)(nil).Info), arg0)
}

// Create mocks base method
func (m *MockEth1Store) Create(ctx context.Context, id string, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, id, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockEth1StoreMockRecorder) Create(ctx, id, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEth1Store)(nil).Create), ctx, id, attr)
}

// Import mocks base method
func (m *MockEth1Store) Import(ctx context.Context, id string, privKey []byte, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, id, privKey, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import
func (mr *MockEth1StoreMockRecorder) Import(ctx, id, privKey, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockEth1Store)(nil).Import), ctx, id, privKey, attr)
}

// Get mocks base method
func (m *MockEth1Store) Get(ctx context.Context, addr common.Address) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockEth1StoreMockRecorder) Get(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEth1Store)(nil).Get), ctx, addr)
}

// List mocks base method
func (m *MockEth1Store) List(ctx context.Context) ([]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockEth1StoreMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEth1Store)(nil).List), ctx)
}

// Update mocks base method
func (m *MockEth1Store) Update(ctx context.Context, addr common.Address, attr *entities.Attributes) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, addr, attr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockEth1StoreMockRecorder) Update(ctx, addr, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEth1Store)(nil).Update), ctx, addr, attr)
}

// Delete mocks base method
func (m *MockEth1Store) Delete(ctx context.Context, addr common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockEth1StoreMockRecorder) Delete(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEth1Store)(nil).Delete), ctx, addr)
}

// GetDeleted mocks base method
func (m *MockEth1Store) GetDeleted(ctx context.Context, addr common.Address) (*entities.ETH1Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleted", ctx, addr)
	ret0, _ := ret[0].(*entities.ETH1Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeleted indicates an expected call of GetDeleted
func (mr *MockEth1StoreMockRecorder) GetDeleted(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleted", reflect.TypeOf((*MockEth1Store)(nil).GetDeleted), ctx, addr)
}

// ListDeleted mocks base method
func (m *MockEth1Store) ListDeleted(ctx context.Context) ([]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeleted", ctx)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeleted indicates an expected call of ListDeleted
func (mr *MockEth1StoreMockRecorder) ListDeleted(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeleted", reflect.TypeOf((*MockEth1Store)(nil).ListDeleted), ctx)
}

// Restore mocks base method
func (m *MockEth1Store) Restore(ctx context.Context, addr common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockEth1StoreMockRecorder) Restore(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockEth1Store)(nil).Restore), ctx, addr)
}

// Destroy mocks base method
func (m *MockEth1Store) Destroy(ctx context.Context, addr common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockEth1StoreMockRecorder) Destroy(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockEth1Store)(nil).Destroy), ctx, addr)
}

// Sign mocks base method
func (m *MockEth1Store) Sign(ctx context.Context, addr common.Address, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockEth1StoreMockRecorder) Sign(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockEth1Store)(nil).Sign), ctx, addr, data)
}

// SignHash mocks base method
func (m *MockEth1Store) SignHash(ctx context.Context, addr common.Address, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignHash", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignHash indicates an expected call of SignHash
func (mr *MockEth1StoreMockRecorder) SignHash(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignHash", reflect.TypeOf((*MockEth1Store)(nil).SignHash), ctx, addr, data)
}

// SignTypedData mocks base method
func (m *MockEth1Store) SignTypedData(ctx context.Context, addr common.Address, typedData *core.TypedData) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTypedData", ctx, addr, typedData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTypedData indicates an expected call of SignTypedData
func (mr *MockEth1StoreMockRecorder) SignTypedData(ctx, addr, typedData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTypedData", reflect.TypeOf((*MockEth1Store)(nil).SignTypedData), ctx, addr, typedData)
}

// SignTransaction mocks base method
func (m *MockEth1Store) SignTransaction(ctx context.Context, addr common.Address, chainID *big.Int, tx *types0.Transaction) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction", ctx, addr, chainID, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction
func (mr *MockEth1StoreMockRecorder) SignTransaction(ctx, addr, chainID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockEth1Store)(nil).SignTransaction), ctx, addr, chainID, tx)
}

// SignEEA mocks base method
func (m *MockEth1Store) SignEEA(ctx context.Context, addr common.Address, chainID *big.Int, tx *types0.Transaction, args *ethereum.PrivateArgs) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignEEA", ctx, addr, chainID, tx, args)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignEEA indicates an expected call of SignEEA
func (mr *MockEth1StoreMockRecorder) SignEEA(ctx, addr, chainID, tx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignEEA", reflect.TypeOf((*MockEth1Store)(nil).SignEEA), ctx, addr, chainID, tx, args)
}

// SignPrivate mocks base method
func (m *MockEth1Store) SignPrivate(ctx context.Context, addr common.Address, tx *types.Transaction) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignPrivate", ctx, addr, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignPrivate indicates an expected call of SignPrivate
func (mr *MockEth1StoreMockRecorder) SignPrivate(ctx, addr, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignPrivate", reflect.TypeOf((*MockEth1Store)(nil).SignPrivate), ctx, addr, tx)
}

// ECRecover mocks base method
func (m *MockEth1Store) ECRecover(ctx context.Context, data, sig []byte) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ECRecover", ctx, data, sig)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ECRecover indicates an expected call of ECRecover
func (mr *MockEth1StoreMockRecorder) ECRecover(ctx, data, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ECRecover", reflect.TypeOf((*MockEth1Store)(nil).ECRecover), ctx, data, sig)
}

// Verify mocks base method
func (m *MockEth1Store) Verify(ctx context.Context, addr common.Address, data, sig []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, addr, data, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockEth1StoreMockRecorder) Verify(ctx, addr, data, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockEth1Store)(nil).Verify), ctx, addr, data, sig)
}

// VerifyTypedData mocks base method
func (m *MockEth1Store) VerifyTypedData(ctx context.Context, addr common.Address, typedData *core.TypedData, sig []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTypedData", ctx, addr, typedData, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTypedData indicates an expected call of VerifyTypedData
func (mr *MockEth1StoreMockRecorder) VerifyTypedData(ctx, addr, typedData, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTypedData", reflect.TypeOf((*MockEth1Store)(nil).VerifyTypedData), ctx, addr, typedData, sig)
}

// Encrypt mocks base method
func (m *MockEth1Store) Encrypt(ctx context.Context, addr common.Address, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockEth1StoreMockRecorder) Encrypt(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEth1Store)(nil).Encrypt), ctx, addr, data)
}

// Decrypt mocks base method
func (m *MockEth1Store) Decrypt(ctx context.Context, addr common.Address, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", ctx, addr, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt
func (mr *MockEth1StoreMockRecorder) Decrypt(ctx, addr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockEth1Store)(nil).Decrypt), ctx, addr, data)
}
