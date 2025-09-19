// Package base provides base schemas and common structures
package base

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
)

// BaseRequest represents the base structure for all API requests
type BaseRequest struct {
	Webhook *string `json:"webhook,omitempty" validate:"omitempty,url"`
	TrackID *string `json:"track_id,omitempty"`
}

// FileInput represents different ways to provide file input
type FileInput struct {
	URL      *string               `json:"url,omitempty" validate:"omitempty,url"`
	Base64   *string               `json:"base64,omitempty"`
	FilePath *string               `json:"file_path,omitempty"`
	File     *multipart.FileHeader `json:"-"`
}

// MarshalJSON implements custom JSON marshaling for FileInput
func (f *FileInput) MarshalJSON() ([]byte, error) {
	if f.URL != nil {
		return json.Marshal(*f.URL)
	}
	if f.Base64 != nil {
		return json.Marshal(*f.Base64)
	}
	if f.FilePath != nil {
		return json.Marshal(*f.FilePath)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements custom JSON unmarshaling for FileInput
func (f *FileInput) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str != "" {
			// Try to determine if it's a URL, base64, or file path
			if len(str) > 7 && (str[:7] == "http://" || str[:8] == "https://") {
				f.URL = &str
			} else {
				f.Base64 = &str
			}
		}
		return nil
	}
	return fmt.Errorf("invalid file input format")
}

// Validate checks if the FileInput has at least one valid input method
func (f *FileInput) Validate() error {
	if f.URL == nil && f.Base64 == nil && f.FilePath == nil && f.File == nil {
		return fmt.Errorf("at least one file input method must be provided")
	}
	return nil
}

// FetchRequest represents a request to fetch results by ID
type FetchRequest struct {
	ID  string `json:"id" validate:"required"`
	Key string `json:"key" validate:"required"`
}

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	ID      string      `json:"id,omitempty"`
}

// EnterpriseRequest provides common enterprise functionality
type EnterpriseRequest struct {
	BaseRequest
}

// SystemDetailsResponse represents system details for enterprise users
type SystemDetailsResponse struct {
	Response
	SystemInfo map[string]interface{} `json:"system_info,omitempty"`
}

// Pointer helper functions for optional fields
func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}
