package aliasent

import (
	"context"
)

//go:generate mockgen -source=backend.go -destination=mock/backend.go -package=mock

// Alias handles the aliases.
type AliasBackend interface {
	// CreateAlias creates an alias in the registry.
	CreateAlias(ctx context.Context, registry RegistryName, alias Alias) (*Alias, error)
	// GetAlias gets an alias from the registry.
	GetAlias(ctx context.Context, registry RegistryName, aliasKey AliasKey) (*Alias, error)
	// UpdateAlias updates an alias in the registry.
	UpdateAlias(ctx context.Context, registry RegistryName, alias Alias) (*Alias, error)
	// GetAlias deletes an alias from the registry.
	DeleteAlias(ctx context.Context, registry RegistryName, aliasKey AliasKey) error

	// ListAlias lists all aliases from a registry.
	ListAliases(ctx context.Context, registry RegistryName) ([]Alias, error)

	// DeleteRegistry deletes a registry, with all the aliases it contained.
	DeleteRegistry(ctx context.Context, registry RegistryName) error
}
