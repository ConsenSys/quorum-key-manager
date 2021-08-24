package eth1

import (
	"context"
	"fmt"
	"testing"

	mock3 "github.com/consensys/quorum-key-manager/src/auth/mock"
	"github.com/consensys/quorum-key-manager/src/auth/types"

	"github.com/consensys/quorum-key-manager/src/infra/log/testutils"
	mock2 "github.com/consensys/quorum-key-manager/src/stores/database/mock"
	testutils2 "github.com/consensys/quorum-key-manager/src/stores/entities/testutils"
	"github.com/consensys/quorum-key-manager/src/stores/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateKey(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	acc := testutils2.FakeETH1Account()
	attributes := testutils2.FakeAttributes()
	key := testutils2.FakeKey()
	key.ID = acc.KeyID
	acc.Tags = attributes.Tags
	expectedErr := fmt.Errorf("error")

	store := mock.NewMockKeyStore(ctrl)
	db := mock2.NewMockETH1Accounts(ctrl)
	logger := testutils.NewMockLogger(ctrl)
	auth := mock3.NewMockAuthorizator(ctrl)

	connector := NewConnector(store, db, auth, logger)

	t.Run("should create eth1Account successfully", func(t *testing.T) {
		auth.EXPECT().Check(&types.Operation{Action: types.ActionWrite, Resource: types.ResourceEth1Account}).Return(nil)
		store.EXPECT().Create(gomock.Any(), key.ID, eth1Algo, attributes).Return(key, nil)
		db.EXPECT().Add(gomock.Any(), newEth1Account(key, attributes)).Return(acc, nil)

		rAcc, err := connector.Create(ctx, key.ID, attributes)

		assert.NoError(t, err)
		assert.Equal(t, rAcc, acc)
	})

	t.Run("should fail with same error if authorization fails", func(t *testing.T) {
		auth.EXPECT().Check(&types.Operation{Action: types.ActionWrite, Resource: types.ResourceEth1Account}).Return(expectedErr)

		_, err := connector.Create(ctx, key.ID, attributes)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to create eth1Account if store fail to create", func(t *testing.T) {
		auth.EXPECT().Check(&types.Operation{Action: types.ActionWrite, Resource: types.ResourceEth1Account}).Return(nil)
		store.EXPECT().Create(gomock.Any(), key.ID, eth1Algo, attributes).Return(nil, expectedErr)

		_, err := connector.Create(ctx, key.ID, attributes)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("should fail to create eth1Account if db fail to add", func(t *testing.T) {
		auth.EXPECT().Check(&types.Operation{Action: types.ActionWrite, Resource: types.ResourceEth1Account}).Return(nil)
		store.EXPECT().Create(gomock.Any(), key.ID, eth1Algo, attributes).Return(key, nil)
		db.EXPECT().Add(gomock.Any(), newEth1Account(key, attributes)).Return(acc, expectedErr)

		_, err := connector.Create(ctx, key.ID, attributes)

		assert.Error(t, err)
		assert.Equal(t, err, expectedErr)
	})
}
