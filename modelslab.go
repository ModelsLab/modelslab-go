// Package modelslab provides the main SDK interface for ModelsLab API
package modelslab

import (
	"github.com/modelslab/modelslab-go/pkg/apis/audio"
	"github.com/modelslab/modelslab-go/pkg/apis/community"
	"github.com/modelslab/modelslab-go/pkg/apis/deepfake"
	"github.com/modelslab/modelslab-go/pkg/apis/image_editing"
	"github.com/modelslab/modelslab-go/pkg/apis/interior"
	"github.com/modelslab/modelslab-go/pkg/apis/realtime"
	"github.com/modelslab/modelslab-go/pkg/apis/threed"
	"github.com/modelslab/modelslab-go/pkg/apis/video"
	"github.com/modelslab/modelslab-go/pkg/client"
)

// ModelsLab represents the main SDK client
type ModelsLab struct {
	client     *client.Client
	enterprise bool

	// API modules
	audio        *audio.API
	video        *video.API
	community    *community.API
	deepfake     *deepfake.API
	imageEditing *image_editing.API
	interior     *interior.API
	realtime     *realtime.API
	threed       *threed.API
}

// New creates a new ModelsLab SDK instance
func New(apiKey string) *ModelsLab {
	c := client.New(apiKey)
	return NewWithClient(c, false)
}

// NewEnterprise creates a new ModelsLab SDK instance for enterprise use
func NewEnterprise(apiKey string) *ModelsLab {
	c := client.New(apiKey)
	return NewWithClient(c, true)
}

// NewWithConfig creates a new ModelsLab SDK instance with custom configuration
func NewWithConfig(config *client.Config) *ModelsLab {
	c := client.NewWithConfig(config)
	return NewWithClient(c, false)
}

// NewEnterpriseWithConfig creates a new enterprise ModelsLab SDK instance with custom configuration
func NewEnterpriseWithConfig(config *client.Config) *ModelsLab {
	c := client.NewWithConfig(config)
	return NewWithClient(c, true)
}

// NewWithClient creates a new ModelsLab SDK instance with a custom client
func NewWithClient(c *client.Client, enterprise bool) *ModelsLab {
	return &ModelsLab{
		client:       c,
		enterprise:   enterprise,
		audio:        audio.New(c, enterprise),
		video:        video.New(c, enterprise),
		community:    community.New(c, enterprise),
		deepfake:     deepfake.New(c, enterprise),
		imageEditing: image_editing.New(c, enterprise),
		interior:     interior.New(c, enterprise),
		realtime:     realtime.New(c, enterprise),
		threed:       threed.New(c, enterprise),
	}
}

// Audio returns the audio API module
func (m *ModelsLab) Audio() *audio.API {
	return m.audio
}

// Video returns the video API module
func (m *ModelsLab) Video() *video.API {
	return m.video
}

// Community returns the community API module
func (m *ModelsLab) Community() *community.API {
	return m.community
}

// DeepFake returns the deepfake API module
func (m *ModelsLab) DeepFake() *deepfake.API {
	return m.deepfake
}

// ImageEditing returns the image editing API module
func (m *ModelsLab) ImageEditing() *image_editing.API {
	return m.imageEditing
}

// Interior returns the interior design API module
func (m *ModelsLab) Interior() *interior.API {
	return m.interior
}

// Realtime returns the realtime API module
func (m *ModelsLab) Realtime() *realtime.API {
	return m.realtime
}

// ThreeD returns the 3D generation API module
func (m *ModelsLab) ThreeD() *threed.API {
	return m.threed
}

// GetClient returns the underlying HTTP client
func (m *ModelsLab) GetClient() *client.Client {
	return m.client
}

// IsEnterprise returns whether this is an enterprise instance
func (m *ModelsLab) IsEnterprise() bool {
	return m.enterprise
}
