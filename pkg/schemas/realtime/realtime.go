// Package realtime provides schemas for realtime API operations
package realtime

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// Text2ImageRequest represents a realtime text-to-image request
type Text2ImageRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
}

// Image2ImageRequest represents a realtime image-to-image request
type Image2ImageRequest struct {
	base.BaseRequest
	Prompt            string         `json:"prompt" validate:"required"`
	InitImage         base.FileInput `json:"init_image" validate:"required"`
	NegativePrompt    *string        `json:"negative_prompt,omitempty"`
	Strength          *float64       `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool          `json:"base64,omitempty"`
	Seed              *int64         `json:"seed,omitempty"`
	GuidanceScale     *float64       `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int           `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	Width             *int           `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int           `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int           `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
}

// InpaintingRequest represents a realtime inpainting request
type InpaintingRequest struct {
	base.BaseRequest
	Prompt            string         `json:"prompt" validate:"required"`
	InitImage         base.FileInput `json:"init_image" validate:"required"`
	MaskImage         base.FileInput `json:"mask_image" validate:"required"`
	NegativePrompt    *string        `json:"negative_prompt,omitempty"`
	Strength          *float64       `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool          `json:"base64,omitempty"`
	Seed              *int64         `json:"seed,omitempty"`
	GuidanceScale     *float64       `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int           `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	Width             *int           `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int           `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int           `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
}

// RealtimeResponse represents a realtime API response
type RealtimeResponse struct {
	base.Response
	Images      []string `json:"images,omitempty"`
	ImageData   []string `json:"image_data,omitempty"`
	ProcessTime int      `json:"process_time,omitempty"`
}
