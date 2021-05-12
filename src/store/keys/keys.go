package keys

import (
	"context"
	"time"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
)

//go:generate mockgen -source=keys.go -destination=mocks/keys.go -package=mocks

type Store interface {
	// Info returns store information
	Info(context.Context) (*entities.StoreInfo, error)

	// Create a new key and stores it
	Create(ctx context.Context, id string, alg *entities.Algorithm, attr *entities.Attributes) (*entities.Key, error)

	// Import an externally created key and stores it
	Import(ctx context.Context, id, privKey string, alg *entities.Algorithm, attr *entities.Attributes) (*entities.Key, error)

	// Get the public part of a stored key.
	Get(ctx context.Context, id, version string) (*entities.Key, error)

	// List keys
	List(ctx context.Context) ([]string, error)

	// Update key tags
	Update(ctx context.Context, id string, attr *entities.Attributes) (*entities.Key, error)

	// Refresh key (create new identical version with different TTL)
	Refresh(ctx context.Context, id string, expirationDate time.Time) error

	// Delete secret not permanently, by using Undelete() the secret can be retrieve
	Delete(ctx context.Context, id string) error

	// GetDeleted keys
	GetDeleted(ctx context.Context, id string) (*entities.Key, error)

	// ListDeleted keys
	ListDeleted(ctx context.Context) ([]string, error)

	// Undelete a previously deleted secret
	Undelete(ctx context.Context, id string) error

	// Destroy secret permanently
	Destroy(ctx context.Context, id string) error

	// Sign from any arbitrary data using the specified key
	Sign(ctx context.Context, id, data, version string) (string, error)

	// Encrypt any arbitrary data using a specified key
	Encrypt(ctx context.Context, id, version string, data string) (string, error)

	// Decrypt a single block of encrypted data.
	Decrypt(ctx context.Context, id, version string, data string) (string, error)
}
