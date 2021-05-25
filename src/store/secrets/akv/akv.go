package akv

import (
	"context"
	"path"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/errors"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/log"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/infra/akv"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/secrets"
)

type Store struct {
	client akv.SecretClient
	logger *log.Logger
}

var _ secrets.Store = &Store{}

func New(client akv.SecretClient, logger *log.Logger) *Store {
	return &Store{
		client: client,
		logger: logger,
	}
}

func (s *Store) Info(context.Context) (*entities.StoreInfo, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) Set(ctx context.Context, id, value string, attr *entities.Attributes) (*entities.Secret, error) {
	logger := s.logger.WithField("id", id)
	res, err := s.client.SetSecret(ctx, id, value, attr.Tags)
	if err != nil {
		logger.Error("failed to set secret")
		return nil, err
	}

	logger.Info("secret was set successfully")
	return parseSecretBundle(&res), nil
}

func (s *Store) Get(ctx context.Context, id, version string) (*entities.Secret, error) {
	logger := s.logger.WithField("id", id)
	res, err := s.client.GetSecret(ctx, id, version)
	if err != nil {
		logger.Error("failed to get secret")
		return nil, err
	}

	logger.Info("secret was retrieved successfully")
	return parseSecretBundle(&res), nil
}

func (s *Store) List(ctx context.Context) ([]string, error) {
	items, err := s.client.GetSecrets(ctx, 0)
	if err != nil {
		s.logger.Error("failed to list secrets")
		return nil, err
	}

	var list = []string{}
	for _, secret := range items {
		// path.Base to only retrieve the secretName instead of https://<vaultName>.vault.azure.net/secrets/<secretName>
		// See listSecrets function in https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/keyvault/examples/go-keyvault-msi-example.go
		list = append(list, path.Base(*secret.ID))
	}

	s.logger.Debug("secrets were listed successfully")
	return list, nil
}

func (s *Store) Delete(ctx context.Context, id string) (*entities.Secret, error) {
	logger := s.logger.WithField("id", id)
	res, err := s.client.DeleteSecret(ctx, id)
	if err != nil {
		logger.Error("failed to delete secret")
		return nil, err
	}

	logger.Info("secret was deleted successfully")
	return parseDeleteSecretBundle(&res), nil
}

func (s *Store) GetDeleted(ctx context.Context, id string) (*entities.Secret, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) ListDeleted(ctx context.Context) ([]string, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) Undelete(ctx context.Context, id string) error {
	return errors.ErrNotImplemented
}

func (s *Store) Destroy(ctx context.Context, id string) error {
	logger := s.logger.WithField("id", id)
	_, err := s.client.PurgeDeletedSecret(ctx, id)
	if err != nil {
		logger.Error("failed to destroy secret")
		return err
	}

	logger.Info("secret was destroyed successfully")
	return nil
}
