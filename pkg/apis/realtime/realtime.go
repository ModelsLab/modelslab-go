// Package realtime provides realtime API operations
package realtime

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	baseSchema "github.com/modelslab/modelslab-go/pkg/schemas/base"
	"github.com/modelslab/modelslab-go/pkg/schemas/realtime"
)

// API provides realtime operations
type API struct {
	*base.BaseAPI
}

// New creates a new realtime API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "realtime"),
	}
}

// TextToImage performs realtime text-to-image generation
func (r *API) TextToImage(ctx context.Context, req *realtime.Text2ImageRequest) (*realtime.RealtimeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := r.GetBaseURL() + "text2img"
	resp, err := r.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("realtime text-to-image request failed: %w", err)
	}

	return &realtime.RealtimeResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// ImageToImage performs realtime image-to-image generation
func (r *API) ImageToImage(ctx context.Context, req *realtime.Image2ImageRequest) (*realtime.RealtimeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := r.GetBaseURL() + "img2img"
	resp, err := r.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("realtime image-to-image request failed: %w", err)
	}

	return &realtime.RealtimeResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// Inpainting performs realtime inpainting
func (r *API) Inpainting(ctx context.Context, req *realtime.InpaintingRequest) (*realtime.RealtimeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := r.GetBaseURL() + "inpaint"
	resp, err := r.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("realtime inpainting request failed: %w", err)
	}

	return &realtime.RealtimeResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}
