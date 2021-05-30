package accounts

import (
	"fmt"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/log"
	manifest "github.com/ConsenSysQuorum/quorum-key-manager/src/services/manifests/types"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/manager/akv"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/manager/hashicorp"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/store/database"
	eth1 "github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/store/eth1/local"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores/types"
)

type Eth1Specs struct {
	Keystore manifest.Kind
	Specs    interface{}
}

func NewEth1(specs *Eth1Specs, eth1Accounts database.ETH1Accounts, logger *log.Logger) (*eth1.Store, error) {
	switch specs.Keystore {
	case types.HashicorpKeys:
		spec := &hashicorp.KeySpecs{}
		if err := manifest.UnmarshalSpecs(specs.Specs, spec); err != nil {
			logger.WithError(err).Error("failed to unmarshal Hashicorp keystore specs")
			return nil, err
		}
		store, err := hashicorp.NewKeyStore(spec, logger)
		if err != nil {
			logger.WithError(err).Error("failed to create new Hashicorp Keystore")
			return nil, err
		}
		return eth1.New(store, eth1Accounts, logger), nil
	case types.AKVKeys:
		spec := &akv.KeySpecs{}
		if err := manifest.UnmarshalSpecs(specs.Specs, spec); err != nil {
			logger.WithError(err).Error("failed to unmarshal AKV keystore specs")
			return nil, err
		}
		store, err := akv.NewKeyStore(spec, logger)
		if err != nil {
			logger.WithError(err).Error("failed to create new AKV Keystore")
			return nil, err
		}
		return eth1.New(store, eth1Accounts, logger), nil
	default:
		err := fmt.Errorf("invalid keystore kind %s", specs.Keystore)
		logger.WithError(err).Error()
		return nil, err
	}
}
