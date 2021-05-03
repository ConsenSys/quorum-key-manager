package client

import (
	"context"
	"net/http"
)

const storeIDHeader = "X-Store-Id"

type HTTPClient struct {
	client *http.Client
	config *Config
}

var _ KeyManagerClient = &HTTPClient{}

func NewHTTPClient(h *http.Client, c *Config) *HTTPClient {
	return &HTTPClient{
		client: h,
		config: c,
	}
}

func withStore(ctx context.Context, storeName string) context.Context {
	return context.WithValue(ctx, RequestHeaderKey, map[string]string{
		storeIDHeader: storeName,
	})
}
