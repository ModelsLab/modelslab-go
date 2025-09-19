// Package deepfake provides deepfake API operations
package deepfake

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/deepfake"
)

// API provides deepfake-related operations
type API struct {
	*base.BaseAPI
}

// New creates a new deepfake API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "deepfake"),
	}
}

// SpecificFaceSwap performs specific face swap
func (d *API) SpecificFaceSwap(ctx context.Context, req *deepfake.SpecificFaceSwapRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := d.GetBaseURL() + "single_face_swap"
	resp, err := d.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("specific face swap request failed: %w", err)
	}

	return resp, nil
}

// MultipleFaceSwap performs multiple face swap
func (d *API) MultipleFaceSwap(ctx context.Context, req *deepfake.MultipleFaceSwapRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := d.GetBaseURL() + "multiple_face_swap"
	resp, err := d.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("multiple face swap request failed: %w", err)
	}

	return resp, nil
}

// MultipleVideoSwap performs multiple video face swap
func (d *API) MultipleVideoSwap(ctx context.Context, req *deepfake.SpecificVideoSwapRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := d.GetBaseURL() + "specific_video_swap"
	resp, err := d.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("multiple video swap request failed: %w", err)
	}

	return resp, nil
}

// SingleVideoSwap performs single video face swap
func (d *API) SingleVideoSwap(ctx context.Context, req *deepfake.SingleVideoSwapRequest) (*client.APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := d.GetBaseURL() + "single_video_swap"
	resp, err := d.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("single video swap request failed: %w", err)
	}

	return resp, nil
}
