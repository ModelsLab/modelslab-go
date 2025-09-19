// Package video provides schemas for video-related API operations
package video

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// Text2VideoRequest represents a text-to-video generation request
type Text2VideoRequest struct {
	base.BaseRequest
	ModelID           string   `json:"model_id" validate:"required"`
	Prompt            string   `json:"prompt" validate:"required"`
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	Seed              *int64   `json:"seed,omitempty"`
	Width             *int     `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int     `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumFrames         *int     `json:"num_frames,omitempty" validate:"omitempty,min=1,max=300"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	FPS               *int     `json:"fps,omitempty" validate:"omitempty,min=1,max=60"`
}

// Image2VideoRequest represents an image-to-video generation request
type Image2VideoRequest struct {
	base.BaseRequest
	ModelID           string          `json:"model_id" validate:"required"`
	Prompt            string          `json:"prompt" validate:"required"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumFrames         *int            `json:"num_frames,omitempty" validate:"omitempty,min=1,max=300"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	FPS               *int            `json:"fps,omitempty" validate:"omitempty,min=1,max=60"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
}

// Text2VideoUltraRequest represents an ultra text-to-video generation request
type Text2VideoUltraRequest struct {
	base.BaseRequest
	Prompt            string   `json:"prompt" validate:"required"`
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	Seed              *int64   `json:"seed,omitempty"`
	Resolution        *string  `json:"resolution,omitempty" validate:"omitempty,oneof=512x512 768x768 1024x1024 1024x576 576x1024"`
	NumFrames         *int     `json:"num_frames,omitempty" validate:"omitempty,min=1,max=300"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	FPS               *int     `json:"fps,omitempty" validate:"omitempty,min=1,max=60"`
	Portrait          *bool    `json:"portrait,omitempty"`
	SampleShift       *int     `json:"sample_shift,omitempty"`
}

// VideoResponse represents a standard video API response
type VideoResponse struct {
	base.Response
	VideoURL  string `json:"video_url,omitempty"`
	VideoData string `json:"video_data,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	FPS       int    `json:"fps,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}
