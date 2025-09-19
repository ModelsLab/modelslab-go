// Package community provides schemas for community API operations
package community

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// Text2ImageRequest represents a text-to-image generation request
type Text2ImageRequest struct {
	base.BaseRequest
	Prompt            string   `json:"prompt" validate:"required"`
	ModelID           *string  `json:"model_id,omitempty"`
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	Width             *int     `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int     `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int     `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	SafetyChecker     *string  `json:"safety_checker,omitempty"`
	Seed              *int64   `json:"seed,omitempty"`
	EnhancePrompt     *string  `json:"enhance_prompt,omitempty"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	MultiLingual      *string  `json:"multi_lingual,omitempty"`
	Panorama          *string  `json:"panorama,omitempty"`
	SelfAttention     *string  `json:"self_attention,omitempty"`
	Upscale           *string  `json:"upscale,omitempty"`
	LoraModel         *string  `json:"lora_model,omitempty"`
	LoraStrength      *float64 `json:"lora_strength,omitempty" validate:"omitempty,min=0,max=2"`
	Scheduler         *string  `json:"scheduler,omitempty"`
	ClipSkip          *int     `json:"clip_skip,omitempty" validate:"omitempty,min=1,max=12"`
}

// Image2ImageRequest represents an image-to-image generation request
type Image2ImageRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	ModelID           *string         `json:"model_id,omitempty"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	SafetyChecker     *string         `json:"safety_checker,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	EnhancePrompt     *string         `json:"enhance_prompt,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	MultiLingual      *string         `json:"multi_lingual,omitempty"`
	Panorama          *string         `json:"panorama,omitempty"`
	SelfAttention     *string         `json:"self_attention,omitempty"`
	Upscale           *string         `json:"upscale,omitempty"`
	LoraModel         *string         `json:"lora_model,omitempty"`
	LoraStrength      *float64        `json:"lora_strength,omitempty" validate:"omitempty,min=0,max=2"`
	Scheduler         *string         `json:"scheduler,omitempty"`
	ClipSkip          *int            `json:"clip_skip,omitempty" validate:"omitempty,min=1,max=12"`
}

// InpaintingRequest represents an inpainting request
type InpaintingRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	ModelID           *string         `json:"model_id,omitempty"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	MaskImage         *base.FileInput `json:"mask_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	SafetyChecker     *string         `json:"safety_checker,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	EnhancePrompt     *string         `json:"enhance_prompt,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	MultiLingual      *string         `json:"multi_lingual,omitempty"`
	Panorama          *string         `json:"panorama,omitempty"`
	SelfAttention     *string         `json:"self_attention,omitempty"`
	Upscale           *string         `json:"upscale,omitempty"`
	LoraModel         *string         `json:"lora_model,omitempty"`
	LoraStrength      *float64        `json:"lora_strength,omitempty" validate:"omitempty,min=0,max=2"`
	Scheduler         *string         `json:"scheduler,omitempty"`
	ClipSkip          *int            `json:"clip_skip,omitempty" validate:"omitempty,min=1,max=12"`
}

// ControlNetRequest represents a ControlNet request
type ControlNetRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	ModelID           *string         `json:"model_id,omitempty"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	MaskImage         *base.FileInput `json:"mask_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	SafetyChecker     *string         `json:"safety_checker,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	EnhancePrompt     *string         `json:"enhance_prompt,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	MultiLingual      *string         `json:"multi_lingual,omitempty"`
	Panorama          *string         `json:"panorama,omitempty"`
	SelfAttention     *string         `json:"self_attention,omitempty"`
	Upscale           *string         `json:"upscale,omitempty"`
	LoraModel         *string         `json:"lora_model,omitempty"`
	LoraStrength      *float64        `json:"lora_strength,omitempty" validate:"omitempty,min=0,max=2"`
	Scheduler         *string         `json:"scheduler,omitempty"`
	ClipSkip          *int            `json:"clip_skip,omitempty" validate:"omitempty,min=1,max=12"`
	ControlNetModel   *string         `json:"controlnet_model,omitempty"`
	ControlNetType    *string         `json:"controlnet_type,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
}

// ImageResponse represents a standard image generation response
type ImageResponse struct {
	base.Response
	Images    []string `json:"images,omitempty"`
	ImageData []string `json:"image_data,omitempty"`
	Seed      int64    `json:"seed,omitempty"`
}
