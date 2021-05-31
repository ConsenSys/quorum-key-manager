package src

import (
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/app"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/server"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/log"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/middleware"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/manifests"
	manifestsmanager "github.com/ConsenSysQuorum/quorum-key-manager/src/services/manifests/manager"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/nodes"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/services/stores"
)

type Config struct {
	HTTP      *server.Config
	Logger    *log.Config
	Manifests *manifestsmanager.Config
}

func New(cfg *Config, logger *log.Logger) (*app.App, error) {
	// Create app
	a := app.New(&app.Config{HTTP: cfg.HTTP}, logger)

	// Register Service Configuration
	err := a.RegisterServiceConfig(cfg.Manifests)
	if err != nil {
		return nil, err
	}

	// Register Services
	err = manifests.RegisterService(a)
	if err != nil {
		return nil, err
	}

	err = stores.RegisterService(a)
	if err != nil {
		return nil, err
	}

	err = nodes.RegisterService(a)
	if err != nil {
		return nil, err
	}

	// Set Middleware
	err = a.SetMiddleware(middleware.AccessLog(cfg.Logger))
	if err != nil {
		return nil, err
	}

	return a, nil
}
