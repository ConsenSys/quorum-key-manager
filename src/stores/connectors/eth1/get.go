package eth1

import (
	"context"

	"github.com/consensys/quorum-key-manager/src/auth/types"

	"github.com/consensys/quorum-key-manager/src/stores/entities"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func (c Connector) Get(ctx context.Context, addr ethcommon.Address) (*entities.ETH1Account, error) {
	logger := c.logger.With("address", addr.Hex())

	err := c.authorizator.CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceEth1Account})
	if err != nil {
		return nil, err
	}

	acc, err := c.db.Get(ctx, addr.Hex())
	if err != nil {
		return nil, err
	}

	logger.Debug("ethereum account retrieved successfully")
	return acc, nil
}

func (c Connector) GetDeleted(ctx context.Context, addr ethcommon.Address) (*entities.ETH1Account, error) {
	logger := c.logger.With("address", addr.Hex())

	err := c.authorizator.CheckPermission(&types.Operation{Action: types.ActionRead, Resource: types.ResourceEth1Account})
	if err != nil {
		return nil, err
	}

	acc, err := c.db.GetDeleted(ctx, addr.Hex())
	if err != nil {
		return nil, err
	}

	logger.Debug("deleted ethereum account retrieved successfully")
	return acc, nil
}
