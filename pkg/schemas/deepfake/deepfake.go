// Package deepfake provides schemas for deepfake API operations
package deepfake

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// SpecificFaceSwapRequest represents a specific face swap request
type SpecificFaceSwapRequest struct {
	base.BaseRequest
	InitImage      base.FileInput `json:"init_image" validate:"required"`
	TargetImage    base.FileInput `json:"target_image" validate:"required"`
	ReferenceImage base.FileInput `json:"reference_image" validate:"required"`
	Watermark      *bool          `json:"watermark,omitempty"`
}

// MultipleFaceSwapRequest represents a multiple face swap request
type MultipleFaceSwapRequest struct {
	base.BaseRequest
	InitImage   base.FileInput `json:"init_image" validate:"required"`
	TargetImage base.FileInput `json:"target_image" validate:"required"`
	Watermark   *bool          `json:"watermark,omitempty"`
}

// SingleVideoSwapRequest represents a single video face swap request
type SingleVideoSwapRequest struct {
	base.BaseRequest
	InitImage    base.FileInput `json:"init_image" validate:"required"`
	InitVideo    base.FileInput `json:"init_video" validate:"required"`
	OutputFormat *string        `json:"output_format,omitempty" validate:"omitempty,oneof=mp4 avi mov"`
	Watermark    *bool          `json:"watermark,omitempty"`
}

// SpecificVideoSwapRequest represents a specific video face swap request
type SpecificVideoSwapRequest struct {
	base.BaseRequest
	InitImage      base.FileInput `json:"init_image" validate:"required"`
	InitVideo      base.FileInput `json:"init_video" validate:"required"`
	ReferenceImage base.FileInput `json:"reference_image" validate:"required"`
	OutputFormat   *string        `json:"output_format,omitempty" validate:"omitempty,oneof=mp4 avi mov"`
	Watermark      *bool          `json:"watermark,omitempty"`
}

// DeepFakeResponse represents a deepfake operation response
type DeepFakeResponse struct {
	base.Response
	ResultURL   string `json:"result_url,omitempty"`
	ResultData  string `json:"result_data,omitempty"`
	ProcessTime int    `json:"process_time,omitempty"`
}
