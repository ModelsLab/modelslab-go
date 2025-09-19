// Package image_editing provides image editing API operations
package image_editing

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/image_editing"
)

// API provides image editing operations
type API struct {
	*base.BaseAPI
}

// New creates a new image editing API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "image_editing"),
	}
}

// Outpainting performs outpainting
func (i *API) Outpainting(ctx context.Context, req *image_editing.OutpaintingRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "outpaint"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("outpainting request failed: %w", err)
	}

	return resp, nil
}

// BackgroundRemover performs background removal
func (i *API) BackgroundRemover(ctx context.Context, req *image_editing.BackgroundRemoverRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "removebg_mask"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("background removal request failed: %w", err)
	}

	return resp, nil
}

// SuperResolution performs super resolution
func (i *API) SuperResolution(ctx context.Context, req *image_editing.SuperResolutionRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "super_resolution"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("super resolution request failed: %w", err)
	}

	return resp, nil
}

// Fashion performs fashion generation
func (i *API) Fashion(ctx context.Context, req *image_editing.FashionRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "fashion"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("fashion request failed: %w", err)
	}

	return resp, nil
}

// ObjectRemover performs object removal
func (i *API) ObjectRemover(ctx context.Context, req *image_editing.ObjectRemovalRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "object_removal"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("object removal request failed: %w", err)
	}

	return resp, nil
}

// FaceGen performs face generation
func (i *API) FaceGen(ctx context.Context, req *image_editing.FacegenRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "face_gen"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("face generation request failed: %w", err)
	}

	return resp, nil
}

// Inpainting performs inpainting
func (i *API) Inpainting(ctx context.Context, req *image_editing.InpaintingRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "inpaint"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("inpainting request failed: %w", err)
	}

	return resp, nil
}

// Headshot performs headshot generation
func (i *API) Headshot(ctx context.Context, req *image_editing.HeadshotRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "head_shot"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("headshot request failed: %w", err)
	}

	return resp, nil
}

// FluxHeadshot performs flux headshot generation
func (i *API) FluxHeadshot(ctx context.Context, req *image_editing.FluxHeadshotRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := i.GetBaseURL() + "flux_headshot"
	resp, err := i.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("flux headshot request failed: %w", err)
	}

	return resp, nil
}
