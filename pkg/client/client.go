// Package client provides the core HTTP client for ModelsLab API
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

// Client represents the ModelsLab API client
type Client struct {
	apiKey       string
	baseURL      string
	httpClient   *http.Client
	fetchRetry   int
	fetchTimeout time.Duration
	validator    *validator.Validate
}

// Config holds configuration options for the client
type Config struct {
	APIKey       string        `validate:"required"`
	BaseURL      string        `validate:"required,url"`
	FetchRetry   int           `validate:"min=1,max=100"`
	FetchTimeout time.Duration `validate:"min=1s,max=300s"`
	HTTPTimeout  time.Duration `validate:"min=1s,max=600s"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL:      "https://modelslab.com/api/",
		FetchRetry:   10,
		FetchTimeout: 2 * time.Second,
		HTTPTimeout:  30 * time.Second,
	}
}

// New creates a new ModelsLab client with the provided API key
func New(apiKey string) *Client {
	config := DefaultConfig()
	config.APIKey = apiKey
	return NewWithConfig(config)
}

// NewWithConfig creates a new ModelsLab client with custom configuration
func NewWithConfig(config *Config) *Client {
	if config.APIKey == "" {
		config.APIKey = os.Getenv("MODELSLAB_API_KEY")
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		panic(fmt.Sprintf("invalid client configuration: %v", err))
	}

	return &Client{
		apiKey:       config.APIKey,
		baseURL:      config.BaseURL,
		fetchRetry:   config.FetchRetry,
		fetchTimeout: config.FetchTimeout,
		validator:    validate,
		httpClient: &http.Client{
			Timeout: config.HTTPTimeout,
		},
	}
}

// APIResponse represents a standard API response
// Changed to map[string]interface{} to preserve all fields from API
type APIResponse map[string]interface{}

// APIError represents an API error
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error %d: %s - %s", e.StatusCode, e.Message, e.Details)
	}
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

func (c *Client) Post(ctx context.Context, endpoint string, data interface{}) (*APIResponse, error) {
	if data != nil {
		if err := c.validator.Struct(data); err != nil {
			return nil, fmt.Errorf("validation error: %w", err)
		}
	}

	requestData := map[string]interface{}{
		"key": c.apiKey,
	}

	if data != nil {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request data: %w", err)
		}

		var dataMap map[string]interface{}
		if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
			return nil, fmt.Errorf("failed to unmarshal request data: %w", err)
		}

		for key, value := range dataMap {
			requestData[key] = value
		}
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "Request failed",
			Details:    string(body),
		}
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// Fetch performs a fetch operation with retry logic
func (c *Client) Fetch(ctx context.Context, endpoint, id string) (*APIResponse, error) {
	fetchURL := fmt.Sprintf("%sfetch/%s", endpoint, id)

	var lastErr error
	for i := 0; i < c.fetchRetry; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		resp, err := c.Post(ctx, fetchURL, map[string]string{"key": c.apiKey})
		if err != nil {
			lastErr = err
			time.Sleep(c.fetchTimeout)
			continue
		}

		if status, ok := (*resp)["status"].(string); ok && status == "success" {
			return resp, nil
		}

		message := ""
		if msg, ok := (*resp)["message"].(string); ok {
			message = msg
		}
		lastErr = fmt.Errorf("fetch failed: %s", message)
		time.Sleep(c.fetchTimeout)
	}

	return nil, fmt.Errorf("fetch failed after %d retries: %w", c.fetchRetry, lastErr)
}

// GetAPIKey returns the configured API key
func (c *Client) GetAPIKey() string {
	return c.apiKey
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// SetHTTPClient allows setting a custom HTTP client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}
