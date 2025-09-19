// Package base provides base functionality for all API modules
package base

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/client"
)

// BaseAPI provides common functionality for all API modules
type BaseAPI struct {
	client     *client.Client
	enterprise bool
	baseURL    string
}

// NewBaseAPI creates a new base API instance
func NewBaseAPI(c *client.Client, enterprise bool, apiPath string) *BaseAPI {
	var baseURL string
	if enterprise {
		baseURL = c.GetBaseURL() + "v1/enterprise/" + apiPath + "/"
	} else {
		baseURL = c.GetBaseURL() + "v6/" + apiPath + "/"
	}

	return &BaseAPI{
		client:     c,
		enterprise: enterprise,
		baseURL:    baseURL,
	}
}

// Fetch performs a fetch operation with the provided ID
func (b *BaseAPI) Fetch(ctx context.Context, id string) (*client.APIResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required for fetch operation")
	}

	resp, err := b.client.Fetch(ctx, b.baseURL, id)
	if err != nil {
		return nil, fmt.Errorf("fetch operation failed: %w", err)
	}

	return resp, nil
}

// SystemDetails returns system details (enterprise only)
func (b *BaseAPI) SystemDetails(ctx context.Context) (*client.APIResponse, error) {
	if !b.enterprise {
		return nil, fmt.Errorf("system details are only available for enterprise users")
	}

	endpoint := b.baseURL + "system_details"
	resp, err := b.client.Post(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("system details request failed: %w", err)
	}

	return resp, nil
}

// RestartServer restarts the server (enterprise only)
func (b *BaseAPI) RestartServer(ctx context.Context) (*client.APIResponse, error) {
	if !b.enterprise {
		return nil, fmt.Errorf("restart server is only available for enterprise users")
	}

	endpoint := b.baseURL + "restart_server"
	resp, err := b.client.Post(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("restart server request failed: %w", err)
	}

	return resp, nil
}

// Update updates the system (enterprise only)
func (b *BaseAPI) Update(ctx context.Context) (*client.APIResponse, error) {
	if !b.enterprise {
		return nil, fmt.Errorf("update is only available for enterprise users")
	}

	endpoint := b.baseURL + "update"
	resp, err := b.client.Post(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("update request failed: %w", err)
	}

	return resp, nil
}

// ClearCache clears the cache (enterprise only)
func (b *BaseAPI) ClearCache(ctx context.Context) (*client.APIResponse, error) {
	if !b.enterprise {
		return nil, fmt.Errorf("clear cache is only available for enterprise users")
	}

	endpoint := b.baseURL + "clear_cache"
	resp, err := b.client.Post(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("clear cache request failed: %w", err)
	}

	return resp, nil
}

// ClearQueue clears the queue (enterprise only)
func (b *BaseAPI) ClearQueue(ctx context.Context) (*client.APIResponse, error) {
	if !b.enterprise {
		return nil, fmt.Errorf("clear queue is only available for enterprise users")
	}

	endpoint := b.baseURL + "clear_queue"
	resp, err := b.client.Post(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("clear queue request failed: %w", err)
	}

	return resp, nil
}

// GetClient returns the underlying client
func (b *BaseAPI) GetClient() *client.Client {
	return b.client
}

// IsEnterprise returns whether this is an enterprise instance
func (b *BaseAPI) IsEnterprise() bool {
	return b.enterprise
}

// GetBaseURL returns the base URL for this API
func (b *BaseAPI) GetBaseURL() string {
	return b.baseURL
}
