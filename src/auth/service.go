package auth

import (
	"net/http"

	manifestreader "github.com/consensys/quorum-key-manager/src/infra/manifests/filesystem"

	"github.com/consensys/quorum-key-manager/src/auth/authenticator/tls"

	apikey "github.com/consensys/quorum-key-manager/src/auth/authenticator/api-key"
	"github.com/consensys/quorum-key-manager/src/auth/authenticator/oidc"
	"github.com/consensys/quorum-key-manager/src/infra/log"

	"github.com/consensys/quorum-key-manager/pkg/app"
	"github.com/consensys/quorum-key-manager/src/auth/authenticator"
	authmanager "github.com/consensys/quorum-key-manager/src/auth/manager"
)

func RegisterService(a *app.App, logger log.Logger) error {
	// Load configuration
	cfg := new(Config)
	err := a.ServiceConfig(cfg)
	if err != nil {
		return err
	}

	manifestReader, err := manifestreader.New(cfg.Manifest)
	if err != nil {
		return err
	}

	// Create and register the stores service
	policyMngr := authmanager.New(manifestReader, logger)
	err = a.RegisterService(policyMngr)
	if err != nil {
		return err
	}

	return nil
}

func Middleware(a *app.App, logger log.Logger) (func(http.Handler) http.Handler, error) {
	// Load configuration
	cfg := new(Config)
	err := a.ServiceConfig(cfg)
	if err != nil {
		return nil, err
	}

	var auths []authenticator.Authenticator
	if cfg.OIDC != nil {
		oidcAuth, err := oidc.NewAuthenticator(cfg.OIDC)
		if err != nil {
			return nil, err
		} else if oidcAuth != nil {
			logger.Info("OIDC Authenticator is enabled")
			auths = append(auths, oidcAuth)
		}
	}
	if cfg.TLS != nil {
		tlsAuth, err := tls.NewAuthenticator(cfg.TLS)
		if err != nil {
			return nil, err
		} else if tlsAuth != nil {
			logger.Info("TLS Authenticator is enabled")
			auths = append(auths, tlsAuth)
		}
	}

	if cfg.APIKEY != nil {
		apikeyAuth, err := apikey.NewAuthenticator(cfg.APIKEY)
		if err != nil {
			return nil, err
		} else if apikeyAuth != nil {
			logger.Info("API-KEY Authenticator is enabled")
			auths = append(auths, apikeyAuth)
		}
	}

	// Create middleware
	mid := authenticator.NewMiddleware(
		logger,
		auths...,
	)

	return mid.Then, nil
}
