package backend

import (
	"github.com/ConsenSysQuorum/quorum-key-manager/core/audit"
	noopauditor "github.com/ConsenSysQuorum/quorum-key-manager/core/audit/noop"
	authmanager "github.com/ConsenSysQuorum/quorum-key-manager/core/auth/manager"
	noopauthmanager "github.com/ConsenSysQuorum/quorum-key-manager/core/auth/manager/noop"
	storemanager "github.com/ConsenSysQuorum/quorum-key-manager/core/store/manager"
	authenticatedmanager "github.com/ConsenSysQuorum/quorum-key-manager/core/store/manager/auth"
	basemanager "github.com/ConsenSysQuorum/quorum-key-manager/core/store/manager/base"
)

// Backend holds internal Key Manager components and
// makes it available for API components
type Backend interface {
	// StoreManager returns the Store Manager
	StoreManager() storemanager.Manager

	// ManifestLoader returns the Manifest Loader
	ManifestLoader() manifest.Loader

	// AuthManager returns the PolicyManager
	AuthManager() authmanager.Manager

	// Auditor returns the PolicyManager
	Auditor() audit.Auditor
}

type BaseBackend struct {
	auditor    audit.Auditor
	authMngr   authmanager.Manager
	storeMnger storemanager.Manager
	loader     manifest.Loader
}

func New() *BaseBackend {
	// Create auditor
	auditor := noopauditor.New()

	// Create store manager
	storeMngr := basemanager.New(auditor)
	storeMngr = authenticatedmanager.Wrap(storeMngr)

	return &BaseBackend{
		auditor:    auditor,
		authMngr:   noopauthmanager.New(),
		storeMnger: storeMngr,
	}
}

func (bckend *BaseBackend) StoreManager() storemanager.Manager {
	return bckend.storeMnger
}

func (bckend *BaseBackend) ManifestLoader() manifest.Loader {
	return bckend.loader
}

func (bckend *BaseBackend) AuthManager() authmanager.Manager {
	return bckend.authMngr
}

func (bckend *BaseBackend) Auditor() audit.Auditor {
	return bckend.auditor
}
