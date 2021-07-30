package aws

import (
	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/infra/aws/client"
	"github.com/consensys/quorum-key-manager/src/infra/log"
	"github.com/consensys/quorum-key-manager/src/stores/store/database"
	"github.com/consensys/quorum-key-manager/src/stores/store/keys/aws"
)

// KeySpecs is the specs format for an AWS Key Vault key store
type KeySpecs struct {
	Region    string `json:"region"`
	AccessID  string `json:"accessID"`
	SecretKey string `json:"secretKey"`
	Debug     bool   `json:"debug"`
}

func NewKeyStore(specs *KeySpecs, db database.Database, logger log.Logger) (*aws.Store, error) {
	cfg := client.NewConfig(specs.Region, specs.AccessID, specs.SecretKey, specs.Debug)
	cli, err := client.NewKmsClient(cfg)
	if err != nil {
		errMessage := "failed to instantiate AWS client (keys)"
		logger.WithError(err).Error(errMessage, "specs", specs)
		return nil, errors.ConfigError(errMessage)
	}

	store := aws.New(cli, db.Keys(), logger)
	return store, nil
}
