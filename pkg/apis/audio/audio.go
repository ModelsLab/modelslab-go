// Package audio provides audio-related API operations
package audio

import (
	"context"
	"fmt"

	"github.com/modelslab/modelslab-go/pkg/apis/base"
	"github.com/modelslab/modelslab-go/pkg/client"
	"github.com/modelslab/modelslab-go/pkg/schemas/audio"
	baseSchema "github.com/modelslab/modelslab-go/pkg/schemas/base"
)

// API provides audio-related operations
type API struct {
	*base.BaseAPI
}

// New creates a new audio API instance
func New(c *client.Client, enterprise bool) *API {
	return &API{
		BaseAPI: base.NewBaseAPI(c, enterprise, "voice"),
	}
}

// TextToAudio performs text-to-audio conversion
func (a *API) TextToAudio(ctx context.Context, req *audio.Text2AudioRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "text_to_audio"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("text-to-audio request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// TextToSpeech performs text-to-speech conversion
func (a *API) TextToSpeech(ctx context.Context, req *audio.Text2SpeechRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "text_to_speech"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("text-to-speech request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// Voice2Voice performs voice-to-voice conversion
func (a *API) Voice2Voice(ctx context.Context, req *audio.Voice2VoiceRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "voice_to_voice"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("voice-to-voice request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// VoiceCover performs voice cover generation
func (a *API) VoiceCover(ctx context.Context, req *audio.VoiceCoverRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "voice_cover"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("voice cover request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// MusicGen performs music generation
func (a *API) MusicGen(ctx context.Context, req *audio.MusicGenRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "music_gen"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("music generation request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// LyricsGen performs lyrics generation
func (a *API) LyricsGen(ctx context.Context, req *audio.LyricsGeneratorRequest) (*audio.LyricsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "lyrics_generator"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("lyrics generation request failed: %w", err)
	}

	return &audio.LyricsResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// SongGenerator performs song generation
func (a *API) SongGenerator(ctx context.Context, req *audio.SongGeneratorRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "song_generator"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("song generation request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// SpeechToText performs speech-to-text conversion
func (a *API) SpeechToText(ctx context.Context, req *audio.Speech2TextRequest) (*audio.TranscriptionResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "speech_to_text"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("speech-to-text request failed: %w", err)
	}

	return &audio.TranscriptionResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}

// SFXGen performs sound effects generation
func (a *API) SFXGen(ctx context.Context, req *audio.SFXRequest) (*audio.AudioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := a.GetBaseURL() + "sfx"
	resp, err := a.GetClient().Post(ctx, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("SFX generation request failed: %w", err)
	}

	return &audio.AudioResponse{
		Response: baseSchema.Response{
			Status:  resp.Status,
			Message: resp.Message,
			Data:    resp.Data,
			Error:   resp.Error,
		},
	}, nil
}
