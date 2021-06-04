// Code generated by MockGen. DO NOT EDIT.
// Source: aws.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/store/entities"
	kms "github.com/aws/aws-sdk-go/service/kms"
	secretsmanager "github.com/aws/aws-sdk-go/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsManagerClient is a mock of SecretsManagerClient interface.
type MockSecretsManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerClientMockRecorder
}

// MockSecretsManagerClientMockRecorder is the mock recorder for MockSecretsManagerClient.
type MockSecretsManagerClientMockRecorder struct {
	mock *MockSecretsManagerClient
}

// NewMockSecretsManagerClient creates a new mock instance.
func NewMockSecretsManagerClient(ctrl *gomock.Controller) *MockSecretsManagerClient {
	mock := &MockSecretsManagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerClient) EXPECT() *MockSecretsManagerClientMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockSecretsManagerClient) CreateSecret(ctx context.Context, id, value string) (*secretsmanager.CreateSecretOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", ctx, id, value)
	ret0, _ := ret[0].(*secretsmanager.CreateSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretsManagerClientMockRecorder) CreateSecret(ctx, id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).CreateSecret), ctx, id, value)
}

// DeleteSecret mocks base method.
func (m *MockSecretsManagerClient) DeleteSecret(ctx context.Context, id string, force bool) (*secretsmanager.DeleteSecretOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", ctx, id, force)
	ret0, _ := ret[0].(*secretsmanager.DeleteSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretsManagerClientMockRecorder) DeleteSecret(ctx, id, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).DeleteSecret), ctx, id, force)
}

// DescribeSecret mocks base method.
func (m *MockSecretsManagerClient) DescribeSecret(ctx context.Context, id string) (*secretsmanager.DescribeSecretOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecret", ctx, id)
	ret0, _ := ret[0].(*secretsmanager.DescribeSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecret indicates an expected call of DescribeSecret.
func (mr *MockSecretsManagerClientMockRecorder) DescribeSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).DescribeSecret), ctx, id)
}

// GetSecret mocks base method.
func (m *MockSecretsManagerClient) GetSecret(ctx context.Context, id, version string) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, id, version)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretsManagerClientMockRecorder) GetSecret(ctx, id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).GetSecret), ctx, id, version)
}

// ListSecrets mocks base method.
func (m *MockSecretsManagerClient) ListSecrets(ctx context.Context, maxResults int64, nextToken string) (*secretsmanager.ListSecretsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", ctx, maxResults, nextToken)
	ret0, _ := ret[0].(*secretsmanager.ListSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsManagerClientMockRecorder) ListSecrets(ctx, maxResults, nextToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsManagerClient)(nil).ListSecrets), ctx, maxResults, nextToken)
}

// PutSecretValue mocks base method.
func (m *MockSecretsManagerClient) PutSecretValue(ctx context.Context, id, value string) (*secretsmanager.PutSecretValueOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSecretValue", ctx, id, value)
	ret0, _ := ret[0].(*secretsmanager.PutSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSecretValue indicates an expected call of PutSecretValue.
func (mr *MockSecretsManagerClientMockRecorder) PutSecretValue(ctx, id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSecretValue", reflect.TypeOf((*MockSecretsManagerClient)(nil).PutSecretValue), ctx, id, value)
}

// RestoreSecret mocks base method.
func (m *MockSecretsManagerClient) RestoreSecret(ctx context.Context, id string) (*secretsmanager.RestoreSecretOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSecret", ctx, id)
	ret0, _ := ret[0].(*secretsmanager.RestoreSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreSecret indicates an expected call of RestoreSecret.
func (mr *MockSecretsManagerClientMockRecorder) RestoreSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).RestoreSecret), ctx, id)
}

// TagSecretResource mocks base method.
func (m *MockSecretsManagerClient) TagSecretResource(ctx context.Context, id string, tags map[string]string) (*secretsmanager.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagSecretResource", ctx, id, tags)
	ret0, _ := ret[0].(*secretsmanager.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagSecretResource indicates an expected call of TagSecretResource.
func (mr *MockSecretsManagerClientMockRecorder) TagSecretResource(ctx, id, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagSecretResource", reflect.TypeOf((*MockSecretsManagerClient)(nil).TagSecretResource), ctx, id, tags)
}

// UpdateSecret mocks base method.
func (m *MockSecretsManagerClient) UpdateSecret(ctx context.Context, id, value, keyID, desc string) (*secretsmanager.UpdateSecretOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", ctx, id, value, keyID, desc)
	ret0, _ := ret[0].(*secretsmanager.UpdateSecretOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretsManagerClientMockRecorder) UpdateSecret(ctx, id, value, keyID, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretsManagerClient)(nil).UpdateSecret), ctx, id, value, keyID, desc)
}

// MockKmsClient is a mock of KmsClient interface.
type MockKmsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKmsClientMockRecorder
}

// MockKmsClientMockRecorder is the mock recorder for MockKmsClient.
type MockKmsClientMockRecorder struct {
	mock *MockKmsClient
}

// NewMockKmsClient creates a new mock instance.
func NewMockKmsClient(ctrl *gomock.Controller) *MockKmsClient {
	mock := &MockKmsClient{ctrl: ctrl}
	mock.recorder = &MockKmsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKmsClient) EXPECT() *MockKmsClientMockRecorder {
	return m.recorder
}

// CreateKey mocks base method.
func (m *MockKmsClient) CreateKey(ctx context.Context, id string, alg *entities.Algorithm, attr *entities.Attributes) (*kms.CreateKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", ctx, id, alg, attr)
	ret0, _ := ret[0].(*kms.CreateKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockKmsClientMockRecorder) CreateKey(ctx, id, alg, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockKmsClient)(nil).CreateKey), ctx, id, alg, attr)
}

// DeleteKey mocks base method.
func (m *MockKmsClient) DeleteKey(ctx context.Context, id string) (*kms.DisableKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", ctx, id)
	ret0, _ := ret[0].(*kms.DisableKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockKmsClientMockRecorder) DeleteKey(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockKmsClient)(nil).DeleteKey), ctx, id)
}

// GetPublicKey mocks base method.
func (m *MockKmsClient) GetPublicKey(ctx context.Context, name string) (*kms.GetPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", ctx, name)
	ret0, _ := ret[0].(*kms.GetPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockKmsClientMockRecorder) GetPublicKey(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockKmsClient)(nil).GetPublicKey), ctx, name)
}

// ListKeys mocks base method.
func (m *MockKmsClient) ListKeys(ctx context.Context, limit int64, marker string) (*kms.ListKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, limit, marker)
	ret0, _ := ret[0].(*kms.ListKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockKmsClientMockRecorder) ListKeys(ctx, limit, marker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKmsClient)(nil).ListKeys), ctx, limit, marker)
}

// ListTags mocks base method.
func (m *MockKmsClient) ListTags(ctx context.Context, id, marker string) (*kms.ListResourceTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, id, marker)
	ret0, _ := ret[0].(*kms.ListResourceTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockKmsClientMockRecorder) ListTags(ctx, id, marker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockKmsClient)(nil).ListTags), ctx, id, marker)
}

// Sign mocks base method.
func (m *MockKmsClient) Sign(ctx context.Context, id string, msg []byte) (*kms.SignOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, id, msg)
	ret0, _ := ret[0].(*kms.SignOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockKmsClientMockRecorder) Sign(ctx, id, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockKmsClient)(nil).Sign), ctx, id, msg)
}

// Verify mocks base method.
func (m *MockKmsClient) Verify(ctx context.Context, id string, msg, signature []byte) (*kms.VerifyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, id, msg, signature)
	ret0, _ := ret[0].(*kms.VerifyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockKmsClientMockRecorder) Verify(ctx, id, msg, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockKmsClient)(nil).Verify), ctx, id, msg, signature)
}
