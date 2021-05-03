package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/errors"
)

const (
	internalErrMsg    = "Internal server error. Please ask an admin for help or try again later"
	internalDepErrMsg = "Failed dependency. Please ask an admin for help or try again later"
)

type ErrorResponse struct {
	Message string `json:"message" example:"error message"`
	Code    uint64 `json:"code,omitempty" example:"24000"`
}

func WriteHTTPErrorResponse(rw http.ResponseWriter, err error) {
	switch {
	case errors.IsAlreadyExistsError(err):
		writeErrorResponse(rw, http.StatusConflict, err)
	case errors.IsNotFoundError(err):
		writeErrorResponse(rw, http.StatusNotFound, err)
	case errors.IsUnauthorizedError(err):
		writeErrorResponse(rw, http.StatusUnauthorized, err)
	case errors.IsInvalidFormatError(err):
		writeErrorResponse(rw, http.StatusBadRequest, err)
	case errors.IsInvalidParameterError(err), errors.IsEncodingError(err):
		writeErrorResponse(rw, http.StatusUnprocessableEntity, err)
	case errors.IsHashicorpVaultConnectionError(err), errors.IsAKVConnectionError(err), errors.IsDependencyFailureError(err):
		writeErrorResponse(rw, http.StatusFailedDependency, errors.DependencyFailureError(internalDepErrMsg))
	case err != nil:
		writeErrorResponse(rw, http.StatusInternalServerError, errors.InternalError(internalErrMsg))
	}
}

func writeErrorResponse(rw http.ResponseWriter, status int, err error) {
	msg, e := json.Marshal(ErrorResponse{Message: err.Error(), Code: errors.FromError(err).GetCode()})
	if e != nil {
		http.Error(rw, e.Error(), status)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	rw.WriteHeader(status)
	_, _ = rw.Write(msg)
}
