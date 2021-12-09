// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	database "github.com/consensys/quorum-key-manager/src/aliases/database"
	"github.com/consensys/quorum-key-manager/src/entities"
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

// Aliases mocks base method
func (m *MockDatabase) Aliases() database.Aliases {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aliases")
	ret0, _ := ret[0].(database.Aliases)
	return ret0
}

// Aliases indicates an expected call of Aliases
func (mr *MockDatabaseMockRecorder) Aliases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aliases", reflect.TypeOf((*MockDatabase)(nil).Aliases))
}

// MockAliases is a mock of Aliases interface
type MockAliases struct {
	ctrl     *gomock.Controller
	recorder *MockAliasesMockRecorder
}

// MockAliasesMockRecorder is the mock recorder for MockAliases
type MockAliasesMockRecorder struct {
	mock *MockAliases
}

// NewMockAliases creates a new mock instance
func NewMockAliases(ctrl *gomock.Controller) *MockAliases {
	mock := &MockAliases{ctrl: ctrl}
	mock.recorder = &MockAliasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAliases) EXPECT() *MockAliasesMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAliases) Create(ctx context.Context, registry string, alias *entities.Alias) (*entities.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, registry, alias)
	ret0, _ := ret[0].(*entities.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAliasesMockRecorder) Create(ctx, registry, alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAliases)(nil).Create), ctx, registry, alias)
}

// Get mocks base method
func (m *MockAliases) Get(ctx context.Context, registry, aliasKey string) (*entities.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, registry, aliasKey)
	ret0, _ := ret[0].(*entities.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAliasesMockRecorder) Get(ctx, registry, aliasKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAliases)(nil).Get), ctx, registry, aliasKey)
}

// Update mocks base method
func (m *MockAliases) Update(ctx context.Context, registry string, alias *entities.Alias) (*entities.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, registry, alias)
	ret0, _ := ret[0].(*entities.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockAliasesMockRecorder) Update(ctx, registry, alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAliases)(nil).Update), ctx, registry, alias)
}

// Delete mocks base method
func (m *MockAliases) Delete(ctx context.Context, registry, aliasKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, registry, aliasKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAliasesMockRecorder) Delete(ctx, registry, aliasKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAliases)(nil).Delete), ctx, registry, aliasKey)
}

// List mocks base method
func (m *MockAliases) List(ctx context.Context, registry string) ([]entities.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, registry)
	ret0, _ := ret[0].([]entities.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAliasesMockRecorder) List(ctx, registry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAliases)(nil).List), ctx, registry)
}

// DeleteRegistry mocks base method
func (m *MockAliases) DeleteRegistry(ctx context.Context, registry string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRegistry", ctx, registry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRegistry indicates an expected call of DeleteRegistry
func (mr *MockAliasesMockRecorder) DeleteRegistry(ctx, registry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegistry", reflect.TypeOf((*MockAliases)(nil).DeleteRegistry), ctx, registry)
}
