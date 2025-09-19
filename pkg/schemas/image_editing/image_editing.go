// Package image_editing provides schemas for image editing API operations
package image_editing

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// OutpaintingRequest represents an outpainting request
type OutpaintingRequest struct {
	base.BaseRequest
	Prompt            string         `json:"prompt" validate:"required"`
	Image             base.FileInput `json:"image" validate:"required"`
	NegativePrompt    *string        `json:"negative_prompt,omitempty"`
	Width             *int           `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int           `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	OverlapWidth      *int           `json:"overlap_width,omitempty" validate:"omitempty,min=0,max=512"`
	NumInferenceSteps *int           `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64       `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	Seed              *int64         `json:"seed,omitempty"`
}

// BackgroundRemoverRequest represents a background removal request
type BackgroundRemoverRequest struct {
	base.BaseRequest
	Image                           base.FileInput `json:"image" validate:"required"`
	AlphaMatting                    *bool          `json:"alpha_matting,omitempty"`
	PostProcessMask                 *bool          `json:"post_process_mask,omitempty"`
	OnlyMask                        *bool          `json:"only_mask,omitempty"`
	InverseMask                     *bool          `json:"inverse_mask,omitempty"`
	Seed                            *int64         `json:"seed,omitempty"`
	Base64                          *bool          `json:"base64,omitempty"`
	AlphaMattingForegroundThreshold *int           `json:"alpha_matting_foreground_threshold,omitempty" validate:"omitempty,min=0,max=255"`
	AlphaMattingBackgroundThreshold *int           `json:"alpha_matting_background_threshold,omitempty" validate:"omitempty,min=0,max=255"`
	AlphaMattingErodeSize           *int           `json:"alpha_matting_erode_size,omitempty" validate:"omitempty,min=0,max=50"`
}

// SuperResolutionRequest represents a super resolution request
type SuperResolutionRequest struct {
	base.BaseRequest
	InitImage   base.FileInput `json:"init_image" validate:"required"`
	ModelID     *string        `json:"model_id,omitempty"`
	Scale       *int           `json:"scale,omitempty" validate:"omitempty,min=1,max=8"`
	FaceEnhance *bool          `json:"face_enhance,omitempty"`
}

// FashionRequest represents a fashion request
type FashionRequest struct {
	base.BaseRequest
	InitImage         base.FileInput  `json:"init_image" validate:"required"`
	Prompt            *string         `json:"prompt,omitempty"`
	NegativePrompt    *string         `json:"negative_prompt,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	ClothImage        *base.FileInput `json:"cloth_image,omitempty"`
	ClothType         *string         `json:"cloth_type,omitempty"`
}

// ObjectRemovalRequest represents an object removal request
type ObjectRemovalRequest struct {
	base.BaseRequest
	InitImage base.FileInput `json:"init_image" validate:"required"`
	MaskImage base.FileInput `json:"mask_image" validate:"required"`
}

// FacegenRequest represents a face generation request
type FacegenRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	FaceImage         *base.FileInput `json:"face_image,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	SScale            *float64        `json:"s_scale,omitempty" validate:"omitempty,min=1,max=30"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	SafetyChecker     *bool           `json:"safety_checker,omitempty"`
	SafetyCheckerType *string         `json:"safety_checker_type,omitempty"`
	Base64            *bool           `json:"base64,omitempty"`
	Style             *string         `json:"style,omitempty"`
}

// InpaintingRequest represents an inpainting request
type InpaintingRequest struct {
	base.BaseRequest
	Prompt            string         `json:"prompt" validate:"required"`
	InitImage         base.FileInput `json:"init_image" validate:"required"`
	MaskImage         base.FileInput `json:"mask_image" validate:"required"`
	NegativePrompt    *string        `json:"negative_prompt,omitempty"`
	Width             *int           `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int           `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumInferenceSteps *int           `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64       `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	Seed              *int64         `json:"seed,omitempty"`
	Samples           *int           `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	Base64            *bool          `json:"base64,omitempty"`
}

// HeadshotRequest represents a headshot generation request
type HeadshotRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	FaceImage         *base.FileInput `json:"face_image,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	SScale            *float64        `json:"s_scale,omitempty" validate:"omitempty,min=1,max=30"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	SafetyChecker     *bool           `json:"safety_checker,omitempty"`
	SafetyCheckerType *string         `json:"safety_checker_type,omitempty"`
	Base64            *bool           `json:"base64,omitempty"`
	Style             *string         `json:"style,omitempty"`
}

// FluxHeadshotRequest represents a flux headshot generation request
type FluxHeadshotRequest struct {
	base.BaseRequest
	Prompt            string          `json:"prompt" validate:"required"`
	FaceImage         *base.FileInput `json:"face_image,omitempty"`
	Width             *int            `json:"width,omitempty" validate:"omitempty,min=64,max=2048"`
	Height            *int            `json:"height,omitempty" validate:"omitempty,min=64,max=2048"`
	NumInferenceSteps *int            `json:"num_inference_steps,omitempty" validate:"omitempty,min=1,max=150"`
	GuidanceScale     *float64        `json:"guidance_scale,omitempty" validate:"omitempty,min=1,max=30"`
	SScale            *float64        `json:"s_scale,omitempty" validate:"omitempty,min=1,max=30"`
	Samples           *int            `json:"samples,omitempty" validate:"omitempty,min=1,max=10"`
	SafetyChecker     *bool           `json:"safety_checker,omitempty"`
	SafetyCheckerType *string         `json:"safety_checker_type,omitempty"`
	Base64            *bool           `json:"base64,omitempty"`
	Style             *string         `json:"style,omitempty"`
}

// ImageEditingResponse represents a standard image editing response
type ImageEditingResponse struct {
	base.Response
	ResultURL   string   `json:"result_url,omitempty"`
	ResultData  string   `json:"result_data,omitempty"`
	Images      []string `json:"images,omitempty"`
	ProcessTime int      `json:"process_time,omitempty"`
}
