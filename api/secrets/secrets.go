package secretsapi

import (
	"net/http"

	"github.com/ConsenSysQuorum/quorum-key-manager/backend"
)

// New creates a http.Handler to be served on /secrets
func New(bcknd backend.Backend) http.Handler {
	// TODO: to be implemented
	return nil
}
