package storemanager

import (
	"context"
	"fmt"
	"sync"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/manifest"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/store-manager/hashicorp"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/types"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/accounts"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/keys"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/secrets"
)

type manager struct {
	mux     sync.RWMutex
	secrets map[string]*storeBundle
	keys    map[string]*storeBundle
	account map[string]*storeBundle
}

type storeBundle struct {
	manifest *manifest.Manifest
	store    interface{}
}

func New() Manager {
	return &manager{
		mux:     sync.RWMutex{},
		secrets: make(map[string]*storeBundle),
		keys:    make(map[string]*storeBundle),
		account: make(map[string]*storeBundle),
	}
}

func (m *manager) Load(ctx context.Context, mnfsts ...*manifest.Manifest) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	for _, mnf := range mnfsts {
		if err := m.load(ctx, mnf); err != nil {
			return err
		}
	}

	return nil
}

func (m *manager) GetSecretStore(_ context.Context, name string) (secrets.Store, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if storeBundle, ok := m.secrets[name]; ok {
		if store, ok := storeBundle.store.(secrets.Store); ok {
			return store, nil
		}
	}

	return nil, fmt.Errorf("secret store not found")
}

func (m *manager) GetKeyStore(_ context.Context, name string) (keys.Store, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if storeBundle, ok := m.keys[name]; ok {
		if store, ok := storeBundle.store.(keys.Store); ok {
			return store, nil
		}
	}

	return nil, fmt.Errorf("keys store not found")
}

func (m *manager) GetAccountStore(_ context.Context, name string) (accounts.Store, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if storeBundle, ok := m.account[name]; ok {
		if store, ok := storeBundle.store.(accounts.Store); ok {
			return store, nil
		}
	}

	return nil, fmt.Errorf("account store not found")
}

func (m *manager) List(_ context.Context, kind types.Kind) ([]string, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	storeNames := []string{}
	switch kind {
	case "":
		storeNames = append(
			append(m.storeNames(m.secrets, kind), m.storeNames(m.keys, kind)...), m.storeNames(m.account, kind)...)
	case types.HashicorpSecrets, types.AKVSecrets, types.KMSSecrets:
		storeNames = m.storeNames(m.secrets, kind)
	case types.AKVKeys, types.HashicorpKeys, types.KMSKeys:
		storeNames = m.storeNames(m.keys, kind)
	}

	return storeNames, nil
}

func (m *manager) load(_ context.Context, mnf *manifest.Manifest) error {
	switch mnf.Kind {
	case types.HashicorpSecrets:
		spec := &hashicorp.SecretSpecs{}
		if err := mnf.UnmarshalSpecs(spec); err != nil {
			return err
		}
		store, err := hashicorp.NewSecretStore(spec)
		if err != nil {
			return err
		}
		m.secrets[mnf.Name] = &storeBundle{manifest: mnf, store: store}
	default:
		return fmt.Errorf("invalid manifest kind %s", mnf.Kind)
	}

	return nil
}

func (m *manager) storeNames(list map[string]*storeBundle, kind types.Kind) []string {
	var storeNames []string
	for k, store := range list {
		if kind == "" || store.manifest.Kind == kind {
			storeNames = append(storeNames, k)
		}
	}

	return storeNames
}
