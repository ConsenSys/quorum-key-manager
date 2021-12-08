package http

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

// TODO: Make accesslog middleware configurable (at least enable/disable)
// TODO: Move to the metrics/monitoring domain when it exists

type AccessLog struct {
	logger io.Writer
}

func NewAccessLog(accessLogger io.Writer) *AccessLog {
	return &AccessLog{
		logger: accessLogger,
	}
}

func (a *AccessLog) Middleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(a.logger, next)
}
