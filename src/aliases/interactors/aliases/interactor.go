package aliases

import (
	"context"
	"regexp"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/aliases"
	"github.com/consensys/quorum-key-manager/src/aliases/entities"
	"github.com/consensys/quorum-key-manager/src/infra/log"
)

// We make sure Connector implements aliases.Interactor
var _ aliases.Interactor = &Interactor{}

// Interactor is the service layer for other service to query.
type Interactor struct {
	db    aliases.Interactor
	regex *regexp.Regexp

	logger log.Logger
}

func NewInteractor(db aliases.Interactor, logger log.Logger) (*Interactor, error) {
	const aliasParseFormat = `{{(?m)(?P<registry>[a-zA-Z0-9-_+]+):(?P<alias>[a-zA-Z0-9-_+]+)}}$`
	regex, err := regexp.Compile(aliasParseFormat)
	if err != nil {
		return nil, errors.ConfigError("bad regexp format '%v': %v", aliasParseFormat, err)
	}
	return &Interactor{
		db:    db,
		regex: regex,

		logger: logger,
	}, nil
}

func (i *Interactor) validateAliasValue(av entities.AliasValue) error {
	switch av.Kind {
	case entities.KindArray:
		_, err := av.Array()
		if err != nil {
			msg := "bad alias array value"
			i.logger.WithError(err).Error(msg)
			return errors.InvalidParameterError(msg)
		}
	case entities.KindString:
		_, err := av.String()
		if err != nil {
			msg := "bad alias string value"
			i.logger.WithError(err).Error(msg)
			return errors.InvalidParameterError(msg)
		}
	default:
		msg := "bad alias value type"
		i.logger.Error(msg)
		return errors.InvalidParameterError(msg)
	}
	return nil
}

func (i *Interactor) CreateAlias(ctx context.Context, registry string, alias entities.Alias) (*entities.Alias, error) {
	logger := i.logger.With(
		"registry_name", registry,
		"alias_key", alias.Key,
	)

	err := i.validateAliasValue(alias.Value)
	if err != nil {
		return nil, err
	}

	a, err := i.db.CreateAlias(ctx, registry, alias)
	if err != nil {
		return nil, err
	}
	logger.Info("alias created successfully")
	return a, nil
}

func (i *Interactor) GetAlias(ctx context.Context, registry, aliasKey string) (*entities.Alias, error) {
	return i.db.GetAlias(ctx, registry, aliasKey)
}

func (i *Interactor) UpdateAlias(ctx context.Context, registry string, alias entities.Alias) (*entities.Alias, error) {
	logger := i.logger.With(
		"registry_name", registry,
		"alias_key", alias.Key,
	)

	err := i.validateAliasValue(alias.Value)
	if err != nil {
		return nil, err
	}

	a, err := i.db.UpdateAlias(ctx, registry, alias)
	if err != nil {
		return nil, err
	}
	logger.Info("alias updated successfully")
	return a, nil
}

func (i *Interactor) DeleteAlias(ctx context.Context, registry, aliasKey string) error {
	logger := i.logger.With(
		"registry_name", registry,
		"alias_key", aliasKey,
	)
	err := i.db.DeleteAlias(ctx, registry, aliasKey)
	if err != nil {
		return err
	}
	logger.Info("alias deleted successfully")
	return nil
}

func (i *Interactor) ListAliases(ctx context.Context, registry string) ([]entities.Alias, error) {
	return i.db.ListAliases(ctx, registry)
}

func (i *Interactor) DeleteRegistry(ctx context.Context, registry string) error {
	logger := i.logger.With(
		"registry_name", registry,
	)
	err := i.db.DeleteRegistry(ctx, registry)
	if err != nil {
		return err
	}
	logger.Info("registry deleted successfully")
	return nil
}

// ParseAlias parses an alias string and returns the registryName and the aliasKey
// as well as if the string isAlias. If the string is not isAlias, we'll consider it
// as a valid key.
func (i *Interactor) ParseAlias(alias string) (regName, aliasKey string, isAlias bool) {
	submatches := i.regex.FindStringSubmatch(alias)
	if len(submatches) < 3 {
		return "", "", false
	}

	regName = submatches[1]
	aliasKey = submatches[2]

	return regName, aliasKey, true
}

// ReplaceAliases replace a slice of potential aliases with a slice having all the aliases replaced by their value.
// It will fail if no aliases can be found.
func (i *Interactor) ReplaceAliases(ctx context.Context, addrs []string) ([]string, error) {
	var values []string
	for _, addr := range addrs {
		regName, aliasKey, isAlias := i.ParseAlias(addr)

		// it is not an alias
		if !isAlias {
			values = append(values, addr)
			continue
		}

		alias, err := i.db.GetAlias(ctx, regName, aliasKey)
		if err != nil {
			return nil, err
		}

		switch alias.Value.Kind {
		case entities.KindArray:
			vals, ok := alias.Value.Value.([]interface{})
			if !ok {
				return nil, errors.InvalidFormatError("bad array format")
			}

			for _, v := range vals {
				str, ok := v.(string)
				if !ok {
					return nil, errors.InvalidFormatError("bad array value type")
				}

				values = append(values, str)
			}
		case entities.KindString:
			values = append(values, alias.Value.Value.(string))
		default:
			return nil, errors.InvalidFormatError("bad value kind")
		}

	}
	return values, nil
}

// ReplaceSimpleAlias replace a potential alias with its first and only value.
// It will fail if no aliases can be found.
func (i *Interactor) ReplaceSimpleAlias(ctx context.Context, addr string) (string, error) {
	alias, err := i.ReplaceAliases(ctx, []string{addr})
	if err != nil {
		return "", err
	}

	if len(alias) != 1 {
		i.logger.WithError(err).Error("wrong alias type")
		return "", errors.EncodingError("alias should only have 1 value")
	}

	return alias[0], nil
}
