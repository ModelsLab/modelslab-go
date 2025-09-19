// Package interior provides interior design API operations
package interior

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/interior"
)

// API provides interior design operations
type API struct {
	*base.BaseAPI
}

// New creates a new interior API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "interior"),
	}
}

// Interior performs interior design generation
func (i *API) Interior(ctx context.Context, req *interior.InteriorRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "make"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("interior design request failed: %w", err)
	}

	return resp, nil
}

// RoomDecorator performs room decoration
func (i *API) RoomDecorator(ctx context.Context, req *interior.RoomDecoratorRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "room_decorator"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("room decoration request failed: %w", err)
	}

	return resp, nil
}

// Floor performs floor planning
func (i *API) Floor(ctx context.Context, req *interior.FloorRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "floor_planning"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("floor planning request failed: %w", err)
	}

	return resp, nil
}

// Scenario performs scenario generation
func (i *API) Scenario(ctx context.Context, req *interior.ScenarioRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "scenario"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("scenario generation request failed: %w", err)
	}

	return resp, nil
}

// ExteriorRestorer performs exterior restoration
func (i *API) ExteriorRestorer(ctx context.Context, req *interior.ExteriorRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "exterior_restorer"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("exterior restoration request failed: %w", err)
	}

	return resp, nil
}
