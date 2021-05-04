package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/errors"
	jsonutils "github.com/ConsenSysQuorum/quorum-key-manager/pkg/json"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/api/formatters"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/api/types"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/core"
	"github.com/gorilla/mux"
)

type KeysHandler struct {
	backend core.Backend
}

// New creates a http.Handler to be served on /keys
func NewKeysHandler(backend core.Backend) *mux.Router {
	h := &KeysHandler{
		backend: backend,
	}

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/").HandlerFunc(h.create)
	router.Methods(http.MethodPost).Path("/import").HandlerFunc(h.importKey)
	router.Methods(http.MethodPost).Path("/{id}/sign").HandlerFunc(h.sign)

	router.Methods(http.MethodGet).Path("/").HandlerFunc(h.list)
	router.Methods(http.MethodGet).Path("/{id}").HandlerFunc(h.getOne)

	router.Methods(http.MethodDelete).Path("/{id}").HandlerFunc(h.destroy)

	return router
}

func (h *KeysHandler) create(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	createKeyRequest := &types.CreateKeyRequest{}
	err := jsonutils.UnmarshalBody(request.Body, createKeyRequest)
	if err != nil {
		WriteHTTPErrorResponse(rw, errors.InvalidFormatError(err.Error()))
		return
	}

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	key, err := keyStore.Create(
		ctx,
		createKeyRequest.ID,
		&entities.Algorithm{
			Type:          entities.KeyType(createKeyRequest.SigningAlgorithm),
			EllipticCurve: entities.Curve(createKeyRequest.Curve),
		},
		&entities.Attributes{
			Tags: createKeyRequest.Tags,
		})
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatKeyResponse(key))
}

func (h *KeysHandler) importKey(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	importKeyRequest := &types.ImportKeyRequest{}
	err := jsonutils.UnmarshalBody(request.Body, importKeyRequest)
	if err != nil {
		WriteHTTPErrorResponse(rw, errors.InvalidFormatError(err.Error()))
		return
	}

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	key, err := keyStore.Import(
		ctx,
		importKeyRequest.ID,
		importKeyRequest.PrivateKey,
		&entities.Algorithm{
			Type:          entities.KeyType(importKeyRequest.SigningAlgorithm),
			EllipticCurve: entities.Curve(importKeyRequest.Curve),
		},
		&entities.Attributes{
			Tags: importKeyRequest.Tags,
		})
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatKeyResponse(key))
}

func (h *KeysHandler) sign(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	signPayloadRequest := &types.SignPayloadRequest{}
	err := jsonutils.UnmarshalBody(request.Body, signPayloadRequest)
	if err != nil {
		WriteHTTPErrorResponse(rw, errors.InvalidFormatError(err.Error()))
		return
	}

	id := mux.Vars(request)["id"]

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	signature, err := keyStore.Sign(ctx, id, signPayloadRequest.Data, signPayloadRequest.Version)
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	_, _ = rw.Write([]byte(signature))
}

func (h *KeysHandler) getOne(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	id := mux.Vars(request)["id"]
	version := request.URL.Query().Get("version")

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	key, err := keyStore.Get(ctx, id, version)
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatKeyResponse(key))
}

func (h *KeysHandler) list(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	ids, err := keyStore.List(ctx)
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(ids)
}

func (h *KeysHandler) destroy(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	id := mux.Vars(request)["id"]

	keyStore, err := h.backend.StoreManager().GetKeyStore(ctx, getStoreName(request))
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	err = keyStore.Destroy(ctx, id)
	if err != nil {
		WriteHTTPErrorResponse(rw, err)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
