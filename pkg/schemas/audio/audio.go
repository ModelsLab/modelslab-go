// Package audio provides schemas for audio-related API operations
package audio

import "github.com/modelslab/modelslab-go/pkg/schemas/base"

// Text2AudioRequest represents a text-to-audio conversion request
type Text2AudioRequest struct {
	base.BaseRequest
	Prompt    string          `json:"prompt" validate:"required"`
	InitAudio *base.FileInput `json:"init_audio,omitempty"`
	VoiceID   *string         `json:"voice_id,omitempty"`
	Language  *string         `json:"language,omitempty"`
	Speed     *float64        `json:"speed,omitempty" validate:"omitempty,min=0.1,max=10"`
	Temp      *bool           `json:"temp,omitempty"`
	Base64    *bool           `json:"base64,omitempty"`
	Stream    *bool           `json:"stream,omitempty"`
}

// Text2SpeechRequest represents a text-to-speech conversion request
type Text2SpeechRequest struct {
	base.BaseRequest
	Prompt       string   `json:"prompt" validate:"required"`
	VoiceID      *string  `json:"voice_id,omitempty"`
	Language     *string  `json:"language,omitempty"`
	Speed        *float64 `json:"speed,omitempty" validate:"omitempty,min=0.1,max=10"`
	OutputFormat *string  `json:"output_format,omitempty" validate:"omitempty,oneof=wav mp3"`
	Emotion      *string  `json:"emotion,omitempty"`
	Temp         *bool    `json:"temp,omitempty"`
}

// Voice2VoiceRequest represents a voice-to-voice conversion request
type Voice2VoiceRequest struct {
	base.BaseRequest
	InitAudio   base.FileInput `json:"init_audio" validate:"required"`
	TargetAudio base.FileInput `json:"target_audio" validate:"required"`
	Base64      *bool          `json:"base64,omitempty"`
	Temp        *bool          `json:"temp,omitempty"`
	Stream      *bool          `json:"stream,omitempty"`
}

// VoiceCoverRequest represents a voice cover request
type VoiceCoverRequest struct {
	base.BaseRequest
	InitAudio              base.FileInput `json:"init_audio" validate:"required"`
	ModelID                *string        `json:"model_id,omitempty"`
	Pitch                  *string        `json:"pitch,omitempty"`
	Algorithm              *string        `json:"algorithm,omitempty"`
	Rate                   *string        `json:"rate,omitempty"`
	Seed                   *string        `json:"seed,omitempty"`
	Emotion                *string        `json:"emotion,omitempty"`
	Speed                  *float64       `json:"speed,omitempty" validate:"omitempty,min=0.1,max=10"`
	Radius                 *float64       `json:"radius,omitempty"`
	Mix                    *float64       `json:"mix,omitempty" validate:"omitempty,min=0,max=1"`
	HopLength              *int           `json:"hop_length,omitempty"`
	Originality            *float64       `json:"originality,omitempty" validate:"omitempty,min=0,max=1"`
	LeadVoiceVolumeDelta   *int           `json:"lead_voice_volume_delta,omitempty"`
	BackupVoiceVolumeDelta *int           `json:"backup_voice_volume_delta,omitempty"`
	InstrumentVolumeDelta  *int           `json:"instrument_volume_delta,omitempty"`
	ReverbSize             *float64       `json:"reverb_size,omitempty" validate:"omitempty,min=0,max=1"`
	Wetness                *float64       `json:"wetness,omitempty" validate:"omitempty,min=0,max=1"`
	Dryness                *float64       `json:"dryness,omitempty" validate:"omitempty,min=0,max=1"`
	Damping                *float64       `json:"damping,omitempty" validate:"omitempty,min=0,max=1"`
	Base64                 *bool          `json:"base64,omitempty"`
	Temp                   *bool          `json:"temp,omitempty"`
}

// MusicGenRequest represents a music generation request
type MusicGenRequest struct {
	base.BaseRequest
	Prompt       string          `json:"prompt" validate:"required"`
	InitAudio    *base.FileInput `json:"init_audio,omitempty"`
	OutputFormat *string         `json:"output_format,omitempty" validate:"omitempty,oneof=wav mp3 flac"`
	Bitrate      *string         `json:"bitrate,omitempty" validate:"omitempty,oneof=128k 192k 320k"`
	Base64       *bool           `json:"base64,omitempty"`
	Temp         *bool           `json:"temp,omitempty"`
	MaxNewToken  *int            `json:"max_new_token,omitempty" validate:"omitempty,min=1"`
	SamplingRate *int            `json:"sampling_rate,omitempty" validate:"omitempty,min=8000,max=96000"`
}

// LyricsGeneratorRequest represents a lyrics generation request
type LyricsGeneratorRequest struct {
	base.BaseRequest
	Prompt string `json:"prompt" validate:"required"`
}

// SongGeneratorRequest represents a song generation request
type SongGeneratorRequest struct {
	base.BaseRequest
	LyricsGeneration *bool           `json:"lyrics_generation,omitempty"`
	InitAudio        *base.FileInput `json:"init_audio,omitempty"`
	Prompt           string          `json:"prompt" validate:"required"`
	Lyrics           *string         `json:"lyrics,omitempty"`
}

// Speech2TextRequest represents a speech-to-text conversion request
type Speech2TextRequest struct {
	base.BaseRequest
	AudioURL       string  `json:"audio_url" validate:"required,url"`
	InputLanguage  *string `json:"input_language,omitempty"`
	TimestampLevel *string `json:"timestamp_level,omitempty" validate:"omitempty,oneof=word sentence"`
}

// SFXRequest represents a sound effects generation request
type SFXRequest struct {
	base.BaseRequest
	Prompt       string  `json:"prompt" validate:"required"`
	Duration     *int    `json:"duration,omitempty" validate:"omitempty,min=1,max=300"`
	OutputFormat *string `json:"output_format,omitempty" validate:"omitempty,oneof=wav mp3 flac"`
	Bitrate      *string `json:"bitrate,omitempty" validate:"omitempty,oneof=128k 192k 320k"`
	Temp         *bool   `json:"temp,omitempty"`
}

// AudioResponse represents a standard audio API response
type AudioResponse struct {
	base.Response
	AudioURL  string `json:"audio_url,omitempty"`
	AudioData string `json:"audio_data,omitempty"`
}

// LyricsResponse represents a lyrics generation response
type LyricsResponse struct {
	base.Response
	Lyrics string `json:"lyrics,omitempty"`
}

// TranscriptionResponse represents a speech-to-text response
type TranscriptionResponse struct {
	base.Response
	Transcription string                 `json:"transcription,omitempty"`
	Timestamps    []TranscriptionSegment `json:"timestamps,omitempty"`
}

// TranscriptionSegment represents a segment of transcribed text with timing
type TranscriptionSegment struct {
	Text      string  `json:"text"`
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
}
