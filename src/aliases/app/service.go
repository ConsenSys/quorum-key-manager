package app

import (
	"github.com/consensys/quorum-key-manager/src/aliases/api/http"
	db "github.com/consensys/quorum-key-manager/src/aliases/database/postgres"
	"github.com/consensys/quorum-key-manager/src/aliases/service/aliases"
	"github.com/consensys/quorum-key-manager/src/aliases/service/registries"
	"github.com/consensys/quorum-key-manager/src/infra/postgres"
	"github.com/gorilla/mux"

	"github.com/consensys/quorum-key-manager/src/infra/log"
)

func RegisterService(router *mux.Router, logger log.Logger, postgresClient postgres.Client) *aliases.Aliases {
	// Data layer
	aliasRepository := db.NewAlias(postgresClient)
	regisryRepository := db.NewRegistry(postgresClient)

	// Business layer
	aliasService := aliases.New(aliasRepository, regisryRepository, logger)
	registryService := registries.New(regisryRepository, logger)

	// Service layer
	http.NewRegistryHandler(registryService).Register(router)
	http.NewAliasHandler(aliasService).Register(router)

	return aliasService
}
