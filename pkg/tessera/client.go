package tessera

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	httpclient "github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/client"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/request"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/response"
)

// Client is a client to Tessera Private Transaction Manager
type Client interface {
	StoreRaw(ctx context.Context, payload []byte, privateFrom string) (string, error)
}

// HTTPClient is a tessera.Client that uses http
type HTTPClient struct {
	client httpclient.Client
}

// NewHTTPClient creates a new HTTPClient
func NewHTTPClient(c httpclient.Client) *HTTPClient {
	return &HTTPClient{
		client: c,
	}
}

type StoreRawRequest struct {
	Payload     string `json:"payload"`
	PrivateFrom string `json:"privateFrom"`
}

type StoreRawResponse struct {
	Key string `json:"key"`
}

func (c *HTTPClient) StoreRaw(ctx context.Context, payload []byte, privateFrom string) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/storeraw", nil)

	err := request.WriteJSON(req, &StoreRawRequest{
		Payload:     base64.StdEncoding.EncodeToString(payload),
		PrivateFrom: privateFrom,
	})
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	msg := new(StoreRawResponse)
	err = response.ReadJSON(resp, msg)
	if err != nil {
		return "", err
	}

	return msg.Key, nil
}

var ErrNotConfigured = fmt.Errorf("tessera not configured")

// NotConfiguredClient is a Tessera Client that always return a tessera not configured error
type NotConfiguredClient struct{}

func (c *NotConfiguredClient) StoreRaw(context.Context, []byte, string) (string, error) {
	return "", ErrNotConfigured
}
