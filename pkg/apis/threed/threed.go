// Package threed provides 3D generation API operations
package threed

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	baseSchema "github.com/modelslab/modelslab-go/pkg/schemas/base"
	"github.com/modelslab/modelslab-go/pkg/schemas/threed"
)

// API provides 3D generation operations
type API struct {
	*base.BaseAPI
}

// New creates a new 3D API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "3d"),
	}
}

// TextTo3D performs text-to-3D generation
func (t *API) TextTo3D(ctx context.Context, req *threed.Text23DRequest) (*threed.ThreeDResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := t.GetBaseURL() + "text_to_3d"
	resp, err := t.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("text-to-3D request failed: %w", err)
	}

	return &threed.ThreeDResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// ImageTo3D performs image-to-3D generation
func (t *API) ImageTo3D(ctx context.Context, req *threed.Image23DRequest) (*threed.ThreeDResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := t.GetBaseURL() + "image_to_3d"
	resp, err := t.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("image-to-3D request failed: %w", err)
	}

	return &threed.ThreeDResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}
