package vaults

import (
	"context"
	"github.com/consensys/quorum-key-manager/src/stores/entities"
)

func (c *Connector) Get(ctx context.Context, name string) (*entities.Vault, error) {
	logger := c.logger.With("name", name)

	vault, err := c.getVault(ctx, name)
	if err != nil {
		return nil, err
	}

	logger.Debug("vault found successfully")
	return vault, nil
}
