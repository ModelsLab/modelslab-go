// Package video provides video-related API operations
package video

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	baseSchema "github.com/modelslab/modelslab-go/pkg/schemas/base"
	"github.com/modelslab/modelslab-go/pkg/schemas/video"
)

// API provides video-related operations
type API struct {
	*base.BaseAPI
}

// New creates a new video API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "video"),
	}
}

// TextToVideo performs text-to-video generation
func (v *API) TextToVideo(ctx context.Context, req *video.Text2VideoRequest) (*video.VideoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := v.GetBaseURL() + "text2video"
	resp, err := v.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("text-to-video request failed: %w", err)
	}

	return &video.VideoResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// ImageToVideo performs image-to-video generation
func (v *API) ImageToVideo(ctx context.Context, req *video.Image2VideoRequest) (*video.VideoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := v.GetBaseURL() + "img2video"
	resp, err := v.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("image-to-video request failed: %w", err)
	}

	return &video.VideoResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// TextToVideoUltra performs ultra text-to-video generation
func (v *API) TextToVideoUltra(ctx context.Context, req *video.Text2VideoUltraRequest) (*video.VideoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := v.GetBaseURL() + "text2video_ultra"
	resp, err := v.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("ultra text-to-video request failed: %w", err)
	}

	return &video.VideoResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}
