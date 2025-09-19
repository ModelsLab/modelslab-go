// Package interior provides schemas for interior design API operations
package interior

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// InteriorRequest represents an interior design request
type InteriorRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
}

// RoomDecoratorRequest represents a room decoration request
type RoomDecoratorRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
}

// FloorRequest represents a floor planning request
type FloorRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
}

// ScenarioRequest represents a scenario generation request
type ScenarioRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	Scenario          *string         `json:"scenario,omitempty"`
}

// ExteriorRequest represents an exterior restoration request
type ExteriorRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
}

// SketchRenderingRequest represents a sketch rendering request
type SketchRenderingRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	InitImage         *base.FileInput `json:"init_image,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Strength          *float64        `json:"strength,omitempty" validate:"omitempty,min=0,max=2"`
	Base64            *bool           `json:"base64,omitempty"`
	Seed              *int64          `json:"seed,omitempty"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
}

// InteriorResponse represents an interior design response
type InteriorResponse struct {
	base.Response
	ResultURL   string `json:"result_url,omitempty"`
	ResultData  string `json:"result_data,omitempty"`
	ProcessTime int    `json:"process_time,omitempty"`
}
