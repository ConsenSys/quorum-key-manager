package localkeys

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/secrets"
	"github.com/ethereum/go-ethereum/crypto"
)

// Store is a keys.Store that uses a secrets.Store to store privateKey values
// Crypto-operations happen in the memory of the KeyStore and are not delegated to any underlying system
type Store struct {
	secrets secrets.Store
}

// New creates a new localkeys.Store
func New(secretStore secrets.Store) *Store {
	return &Store{
		secrets: secretStore,
	}
}

// Create a new key and stores it
func (s *Store) Create(ctx context.Context, id string, alg *entities.Algo, attr *entities.Attributes) (*entities.Key, error) {
	switch alg.Type {
	case "ecdsa":
		// Generate key
		privKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		// Transform public key into byte
		pubKey := crypto.FromECDSAPub(privKey.Public().(*ecdsa.PublicKey))

		// Set key on the private store
		// TODO: pubkey could be stored as a metadata so we do not need to recompute it each time
		secret, err := s.secrets.Set(ctx, id, string(crypto.FromECDSA(privKey)), attr)
		if err != nil {
			return nil, err
		}

		return &entities.Key{
			PublicKey: pubKey,
			Alg:       alg,
			Attr: &entities.Attributes{
				Tags: secret.Tags,
			},
			Metadata: secret.Metadata,
		}, nil
	default:
		return nil, fmt.Errorf("not supported")
	}

}

// Sign from a digest using the specified key
func (s *Store) Sign(ctx context.Context, id string, data []byte, version string) ([]byte, error) {
	// Get secret value from secret store
	secret, err := s.secrets.Get(ctx, id, version)
	if err != nil {
		return nil, err
	}

	// Mount secret into a private key
	privKey, err := crypto.ToECDSA([]byte(secret.Value))
	if err != nil {
		return nil, err
	}

	// Signs
	return crypto.Sign(data, privKey)
}
