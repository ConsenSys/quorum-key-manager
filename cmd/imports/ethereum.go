package imports

import (
	"context"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/pkg/json"
	"github.com/consensys/quorum-key-manager/src/infra/log"
	manifest "github.com/consensys/quorum-key-manager/src/infra/manifests/entities"
	"github.com/consensys/quorum-key-manager/src/stores"
	"github.com/consensys/quorum-key-manager/src/stores/database"
	"github.com/consensys/quorum-key-manager/src/stores/database/models"
	"github.com/consensys/quorum-key-manager/src/stores/entities"
	eth "github.com/consensys/quorum-key-manager/src/stores/manager/ethereum"
)

func ImportEthereum(ctx context.Context, db database.ETHAccounts, mnf *manifest.Manifest, logger log.Logger) error {
	logger.Info("importing ethereum accounts...")

	store, err := getEthStore(mnf, logger)
	if err != nil {
		return err
	}

	storeIDs, err := store.List(ctx, 0, 0)
	if err != nil {
		return err
	}

	dbAddresses, err := db.SearchAddresses(ctx, false, 0, 0)
	if err != nil {
		return err
	}
	addressMap := arrToMap(dbAddresses)

	var nSuccesses uint
	var nFailures uint
	for _, id := range storeIDs {
		key, err := store.Get(ctx, id)
		if err != nil {
			nFailures++
			continue
		}

		if !key.IsETHAccount() {
			continue
		}

		acc := models.NewETHAccountFromKey(key, &entities.Attributes{})
		if _, found := addressMap[acc.Address.Hex()]; !found {
			_, err = db.Add(ctx, acc)
			if err != nil {
				nFailures++
				continue
			}

			nSuccesses++
		}
	}

	logger.Info("ethereum accounts import completed", "n_successes", nSuccesses, "n_failures", nFailures)
	return nil
}

func getEthStore(mnf *manifest.Manifest, logger log.Logger) (stores.KeyStore, error) {
	if mnf.Kind == manifest.Ethereum {
		spec := &entities.LocalEthSpecs{}
		if err := json.UnmarshalJSON(mnf.Specs, spec); err != nil {
			errMessage := "invalid ethereum store specs"
			logger.WithError(err).Error(errMessage)
			return nil, errors.InvalidFormatError(errMessage)
		}

		return eth.NewLocalEth(spec, nil, logger) // DB here is nil and not the DB we instantiate for the import
	}

	errMessage := "invalid manifest kind for ethereum store"
	logger.Error(errMessage)
	return nil, errors.InvalidFormatError(errMessage)
}
