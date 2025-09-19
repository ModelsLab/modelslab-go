// Package community provides community API operations
package community

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/community"
)

// API provides community-related operations
type API struct {
	*base.BaseAPI
}

// New creates a new community API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "images"),
	}
}

// TextToImage performs text-to-image generation
func (c *API) TextToImage(ctx context.Context, req *community.Text2ImageRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := c.GetBaseURL() + "text2img"
	resp, err := c.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("text-to-image request failed: %w", err)
	}

	return resp, nil
}

// ImageToImage performs image-to-image generation
func (c *API) ImageToImage(ctx context.Context, req *community.Image2ImageRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := c.GetBaseURL() + "img2img"
	resp, err := c.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("image-to-image request failed: %w", err)
	}

	return resp, nil
}

// Inpainting performs inpainting
func (c *API) Inpainting(ctx context.Context, req *community.InpaintingRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := c.GetBaseURL() + "inpaint"
	resp, err := c.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("inpainting request failed: %w", err)
	}

	return resp, nil
}

// ControlNet performs ControlNet generation
func (c *API) ControlNet(ctx context.Context, req *community.ControlNetRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := c.GetBaseURL() + "controlnet"
	resp, err := c.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("ControlNet request failed: %w", err)
	}

	return resp, nil
}
