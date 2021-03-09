package secrets

import (
	"bou.ke/monkey"
	"context"
	"fmt"
	"github.com/ConsenSysQuorum/quorum-key-manager/core/store/models/testutils"
	"github.com/ConsenSysQuorum/quorum-key-manager/infra/vault/mocks"
	"github.com/golang/mock/gomock"
	hashicorp "github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type hashicorpSecretStoreTestSuite struct {
	suite.Suite
	mockVault   *mocks.MockHashicorpVaultClient
	mountPoint  string
	secretStore *hashicorpSecretStore
}

func TestHashicorpSecretStore(t *testing.T) {
	s := new(hashicorpSecretStoreTestSuite)
	suite.Run(t, s)
}

func (s *hashicorpSecretStoreTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.mountPoint = "secret"
	s.mockVault = mocks.NewMockHashicorpVaultClient(ctrl)

	s.secretStore = New(s.mockVault, s.mountPoint)
}

func (s *hashicorpSecretStoreTestSuite) TestSet() {
	ctx := context.Background()
	id := "my-secret"
	value := "my-value"
	expectedPath := s.mountPoint + "/data/" + id
	attributes := testutils.FakeAttributes()
	expectedData := map[string]interface{}{
		valueLabel: value,
		tagsLabel:  attributes.Tags,
	}
	hashicorpSecret := &hashicorp.Secret{
		Data: map[string]interface{}{
			"created_time":  "2018-03-22T02:24:06.945319214Z",
			"deletion_time": "",
			"destroyed":     false,
			"version":       2,
		},
	}

	s.T().Run("should set a new secret successfully", func(t *testing.T) {
		expectedCreatedAt, _ := time.Parse(time.RFC3339, "2018-03-22T02:24:06.945319214Z")

		s.mockVault.EXPECT().Write(expectedPath, expectedData).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Set(ctx, id, value, attributes)

		assert.NoError(t, err)
		assert.Equal(t, value, secret.Value)
		assert.False(t, secret.Disabled)
		assert.Equal(t, expectedCreatedAt, secret.CreatedAt)
		assert.Equal(t, attributes.Tags, secret.Tags)
		assert.Equal(t, 2, secret.Version)
		assert.True(t, secret.ExpireAt.IsZero())
		assert.True(t, secret.DeletedAt.IsZero())
		assert.Nil(t, secret.Recovery)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with same error if write fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("error")

		s.mockVault.EXPECT().Write(s.mountPoint+"/data/"+id, expectedData).Return(nil, expectedErr)

		secret, err := s.secretStore.Set(ctx, id, value, attributes)

		assert.Nil(t, secret)
		assert.Equal(t, expectedErr, err)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with error if it fails to extract metadata", func(t *testing.T) {
		hashSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				"created_time": "invalidTime",
				"version":      2,
			},
		}

		s.mockVault.EXPECT().Write(s.mountPoint+"/data/"+id, expectedData).Return(hashSecret, nil)

		secret, err := s.secretStore.Set(ctx, id, value, attributes)

		assert.Nil(t, secret)
		assert.Error(t, err)
	})
}

func (s *hashicorpSecretStoreTestSuite) TestGet() {
	ctx := context.Background()
	id := "my-secret"
	value := "my-value"
	expectedPath := s.mountPoint + "/data/" + id
	attributes := testutils.FakeAttributes()
	expectedData := map[string]interface{}{
		valueLabel: value,
		tagsLabel:  attributes.Tags,
	}

	s.T().Run("should get a secret successfully with empty version", func(t *testing.T) {
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time":  "2018-03-22T02:24:06.945319214Z",
					"deletion_time": "",
					"destroyed":     false,
					"version":       2,
				},
			},
		}
		expectedCreatedAt, _ := time.Parse(time.RFC3339, "2018-03-22T02:24:06.945319214Z")

		s.mockVault.EXPECT().Read(expectedPath).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, 0)

		assert.NoError(t, err)
		assert.Equal(t, value, secret.Value)
		assert.False(t, secret.Disabled)
		assert.Equal(t, expectedCreatedAt, secret.CreatedAt)
		assert.Equal(t, attributes.Tags, secret.Tags)
		assert.Equal(t, 2, secret.Version)
		assert.True(t, secret.ExpireAt.IsZero())
		assert.True(t, secret.DeletedAt.IsZero())
		assert.Nil(t, secret.Recovery)
	})

	s.T().Run("should get a secret successfully with version", func(t *testing.T) {
		version := 2
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time":  "2018-03-22T02:24:06.945319214Z",
					"deletion_time": "",
					"destroyed":     false,
					"version":       version,
				},
			},
		}

		s.mockVault.EXPECT().Read(fmt.Sprintf("%v?version=%v", expectedPath, version)).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, version)

		assert.NoError(t, err)
		assert.NotNil(t, secret)
	})

	s.T().Run("should get a secret successfully with future deletion time", func(t *testing.T) {
		version := 2
		deletionTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time":  "2018-03-22T02:24:06.945319214Z",
					"deletion_time": deletionTime,
					"destroyed":     false,
					"version":       version,
				},
			},
		}
		expectedExpireAt, _ := time.Parse(time.RFC3339, deletionTime)

		s.mockVault.EXPECT().Read(fmt.Sprintf("%v?version=%v", expectedPath, version)).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, version)

		assert.NoError(t, err)
		assert.Equal(t, expectedExpireAt, secret.ExpireAt)
	})

	s.T().Run("should get a secret successfully with past deletion time", func(t *testing.T) {
		version := 2
		deletionTime := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time":  "2018-03-22T02:24:06.945319214Z",
					"deletion_time": deletionTime,
					"destroyed":     false,
					"version":       version,
				},
			},
		}
		expectedDeletedAt, _ := time.Parse(time.RFC3339, deletionTime)

		s.mockVault.EXPECT().Read(fmt.Sprintf("%v?version=%v", expectedPath, version)).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, version)

		assert.NoError(t, err)
		assert.Empty(t, secret.ExpireAt)
		assert.Equal(t, expectedDeletedAt, secret.DeletedAt)
		assert.True(t, secret.Disabled)
	})

	s.T().Run("should get a secret successfully with past deletion time and destroyed", func(t *testing.T) {
		version := 2
		deletionTime := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time":  "2018-03-22T02:24:06.945319214Z",
					"deletion_time": deletionTime,
					"destroyed":     true,
					"version":       version,
				},
			},
		}
		expectedDeletedAt, _ := time.Parse(time.RFC3339, deletionTime)

		s.mockVault.EXPECT().Read(fmt.Sprintf("%v?version=%v", expectedPath, version)).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, version)

		assert.NoError(t, err)
		assert.Empty(t, secret.ExpireAt)
		assert.Equal(t, expectedDeletedAt, secret.DeletedAt)
		assert.True(t, secret.Disabled)
		assert.Equal(t, expectedDeletedAt, secret.DestroyedAt)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with same error if read fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("error")

		s.mockVault.EXPECT().Read(expectedPath).Return(nil, expectedErr)

		secret, err := s.secretStore.Get(ctx, id, 0)

		assert.Nil(t, secret)
		assert.Equal(t, expectedErr, err)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with error if it fails to extract metadata", func(t *testing.T) {
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				dataLabel: expectedData,
				metadataLabel: map[string]interface{}{
					"created_time": "invalidCreatedTime",
					"version":      1,
				},
			},
		}

		s.mockVault.EXPECT().Read(expectedPath).Return(hashicorpSecret, nil)

		secret, err := s.secretStore.Get(ctx, id, 0)

		assert.Nil(t, secret)
		assert.Error(t, err)
	})
}

func (s *hashicorpSecretStoreTestSuite) TestList() {
	ctx := context.Background()
	expectedPath := s.mountPoint + "/metadata"
	keys := []string{"my-secret1", "my-secret2"}

	s.T().Run("should list all secret ids successfully", func(t *testing.T) {
		hashicorpSecret := &hashicorp.Secret{
			Data: map[string]interface{}{
				"keys": keys,
			},
		}

		s.mockVault.EXPECT().List(expectedPath).Return(hashicorpSecret, nil)

		ids, err := s.secretStore.List(ctx)

		assert.NoError(t, err)
		assert.Equal(t, keys, ids)
	})

	s.T().Run("should return empty list if result is nil", func(t *testing.T) {
		s.mockVault.EXPECT().List(expectedPath).Return(nil, nil)

		ids, err := s.secretStore.List(ctx)

		assert.NoError(t, err)
		assert.Empty(t, ids)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with same error if list fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("error")

		s.mockVault.EXPECT().List(expectedPath).Return(nil, expectedErr)

		secret, err := s.secretStore.List(ctx)

		assert.Nil(t, secret)
		assert.Equal(t, expectedErr, err)
	})
}

func (s *hashicorpSecretStoreTestSuite) TestRefresh() {
	ctx := context.Background()
	id := "my-secret"
	expectedPath := s.mountPoint + "/metadata/" + id
	hashicorpSecret := &hashicorp.Secret{
		Data: map[string]interface{}{},
	}

	s.T().Run("should refresh a secret without expiration date", func(t *testing.T) {
		s.mockVault.EXPECT().Write(expectedPath, map[string]interface{}{}).Return(hashicorpSecret, nil)

		err := s.secretStore.Refresh(ctx, id, time.Time{})

		assert.NoError(t, err)
	})

	s.T().Run("should refresh a secret with expiration date", func(t *testing.T) {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
		})
		expirationDate := time.Date(2021, 1, 1, 2, 0, 0, 0, time.UTC)

		s.mockVault.EXPECT().Write(expectedPath, map[string]interface{}{
			deleteAfterLabel: "1h0m0s",
		}).Return(hashicorpSecret, nil)

		err := s.secretStore.Refresh(ctx, id, expirationDate)

		assert.NoError(t, err)
	})

	// TODO: Implement specific error types and check that the function return the right error type
	s.T().Run("should fail with same error if Write fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("error")

		s.mockVault.EXPECT().Write(expectedPath, map[string]interface{}{}).Return(nil, expectedErr)

		err := s.secretStore.Refresh(ctx, id, time.Time{})

		assert.Equal(t, expectedErr, err)
	})
}
