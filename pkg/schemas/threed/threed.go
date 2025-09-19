// Package threed provides schemas for 3D generation API operations
package threed

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// Text23DRequest represents a text-to-3D generation request
type Text23DRequest struct {
	base.BaseRequest
	Prompt               string   `json:"prompt" validate:"required"`
	Resolution           *string  `json:"resolution,omitempty"`
	OutputFormat         *string  `json:"output_format,omitempty" validate:"omitempty,oneof=obj ply glb"`
	Render               *bool    `json:"render,omitempty"`
	NegativePrompt       *string  `json:"negative_prompt,omitempty"`
	Seed                 *int64   `json:"seed,omitempty"`
	GuidanceScale        *float64 `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	NumInferenceSteps    *int     `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	SSGuidanceStrength   *float64 `json:"ss_guidance_strength,omitempty" validate:"omitempty,min=0,max=10"`
	SlatGuidanceStrength *float64 `json:"slat_guidance_strength,omitempty" validate:"omitempty,min=0,max=10"`
	SlatSamplingSteps    *int     `json:"slat_sampling_steps,omitempty" validate:"omitempty,min=1,max=100"`
	MeshSimplify         *float64 `json:"mesh_simplify,omitempty" validate:"omitempty,min=0,max=1"`
	ForegroundRatio      *float64 `json:"foreground_ratio,omitempty" validate:"omitempty,min=0,max=1"`
	RemoveBG             *bool    `json:"remove_bg,omitempty"`
	ChunkSize            *int     `json:"chunk_size,omitempty" validate:"omitempty,min=1"`
	Temp                 *bool    `json:"temp,omitempty"`
}

// Image23DRequest represents an image-to-3D generation request
type Image23DRequest struct {
	base.BaseRequest
	Image                base.FileInput `json:"image" validate:"required"`
	Resolution           *int           `json:"resolution,omitempty" validate:"omitempty,min=64,max=2048"`
	OutputFormat         *string        `json:"output_format,omitempty" validate:"omitempty,oneof=obj ply glb"`
	Render               *bool          `json:"render,omitempty"`
	Seed                 *int64         `json:"seed,omitempty"`
	MultiImage           *bool          `json:"multi_image,omitempty"`
	SSGuidanceStrength   *float64       `json:"ss_guidance_strength,omitempty" validate:"omitempty,min=0,max=10"`
	SlatGuidanceStrength *float64       `json:"slat_guidance_strength,omitempty" validate:"omitempty,min=0,max=10"`
	SlatSamplingSteps    *int           `json:"slat_sampling_steps,omitempty" validate:"omitempty,min=1,max=100"`
	MeshSimplify         *float64       `json:"mesh_simplify,omitempty" validate:"omitempty,min=0,max=1"`
	ChunkSize            *int           `json:"chunk_size,omitempty" validate:"omitempty,min=1"`
	Temp                 *bool          `json:"temp,omitempty"`
}

// ThreeDResponse represents a 3D generation response
type ThreeDResponse struct {
	base.Response
	ModelURL    string   `json:"model_url,omitempty"`
	ModelData   string   `json:"model_data,omitempty"`
	RenderURL   string   `json:"render_url,omitempty"`
	RenderData  string   `json:"render_data,omitempty"`
	Meshes      []string `json:"meshes,omitempty"`
	ProcessTime int      `json:"process_time,omitempty"`
}
