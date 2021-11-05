package postgres

import (
	"context"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/aliases/database"
	"github.com/consensys/quorum-key-manager/src/aliases/database/models"
	"github.com/consensys/quorum-key-manager/src/aliases/entities"
	"github.com/consensys/quorum-key-manager/src/infra/log"
	"github.com/consensys/quorum-key-manager/src/infra/postgres"
)

var _ database.AliasRepository = &Alias{}

// Alias stores the alias data in a postgres DB.
type Alias struct {
	pgClient postgres.Client
	logger   log.Logger
}

func NewAlias(pgClient postgres.Client, logger log.Logger) *Alias {
	return &Alias{
		pgClient: pgClient,
		logger:   logger,
	}
}

func (s *Alias) CreateAlias(ctx context.Context, registry string, alias entities.Alias) (*entities.Alias, error) {
	logger := s.logger.With(
		"registry_name", registry,
		"alias_key", alias.Key,
	)
	a := models.AliasFromEntity(alias)
	a.RegistryName = registry

	err := s.pgClient.Insert(ctx, &a)
	if err != nil {
		msg := "failed to create alias"
		logger.WithError(err).Error(msg)
		return nil, err
	}
	return &alias, nil
}

func (s *Alias) GetAlias(ctx context.Context, registry, aliasKey string) (*entities.Alias, error) {
	logger := s.logger.With(
		"registry_name", registry,
		"alias_key", aliasKey,
	)
	a := models.Alias{
		Key:          aliasKey,
		RegistryName: registry,
	}
	err := s.pgClient.SelectPK(ctx, &a)
	if err != nil {
		msg := "failed to get alias"
		logger.WithError(err).Error(msg)
		return nil, err
	}

	a.Value, err = s.convertAliasValueFromDB(a.Value)
	if err != nil {
		msg := "failed to convert alias value"
		logger.WithError(err).Error(msg)
		return nil, err
	}

	return a.ToEntity(), nil
}

func (s *Alias) UpdateAlias(ctx context.Context, registry string, alias entities.Alias) (*entities.Alias, error) {
	logger := s.logger.With(
		"registry_name", registry,
		"alias_key", alias.Key,
	)
	a := models.AliasFromEntity(alias)
	a.RegistryName = registry

	err := s.pgClient.UpdatePK(ctx, &a)
	if err != nil {
		msg := "failed to update alias"
		logger.WithError(err).Error(msg)
		return nil, err
	}
	return a.ToEntity(), nil
}

func (s *Alias) DeleteAlias(ctx context.Context, registry, aliasKey string) error {
	logger := s.logger.With(
		"registry_name", registry,
		"alias_key", aliasKey,
	)
	a := models.Alias{
		Key:          aliasKey,
		RegistryName: registry,
	}

	err := s.pgClient.DeletePK(ctx, &a)
	if err != nil {
		msg := "failed to delete alias"
		logger.WithError(err).Error(msg)
		return err
	}
	return nil
}

func (s *Alias) ListAliases(ctx context.Context, registry string) ([]entities.Alias, error) {
	logger := s.logger.With(
		"registry_name", registry,
	)
	reg := registry

	var als []models.Alias
	err := s.pgClient.SelectWhere(ctx, &als, "alias.registry_name = ?", reg)
	if err != nil {
		msg := "failed to list aliases"
		logger.WithError(err).Error(msg)
		return nil, err
	}
	for i := range als {
		als[i].Value, err = s.convertAliasValueFromDB(als[i].Value)
		if err != nil {
			msg := "failed to convert alias value"
			logger.WithError(err).Error(msg)
			return nil, err
		}
	}

	return models.AliasesToEntity(als), nil
}

func (s *Alias) DeleteRegistry(ctx context.Context, registry string) error {
	return errors.NotImplementedError("DeleteRegistry not implemented")
}

func (s *Alias) convertAliasValueFromDB(a entities.AliasValue) (entities.AliasValue, error) {
	switch a.Kind {
	case entities.KindArray:
		// go-pg gets the stored []string as []interface{}, we need to convert
		stored, ok := a.Value.([]interface{})
		if !ok {
			return a, errors.EncodingError("failed to get correct alias array value")
		}
		var s []string
		for _, v := range stored {
			val, ok := v.(string)
			if !ok {
				return a, errors.EncodingError("failed to get correct alias array item value")
			}
			s = append(s, val)
		}
		a.Value = s
	case entities.KindString:
		// we don't need anything specific for KindString as the go-pg unmarshaller converts to the right type.
	default:
		return a, errors.EncodingError("failed to get correct alias value")
	}
	return a, nil
}
