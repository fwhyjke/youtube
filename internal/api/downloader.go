package api

import (
	"context"
	"io"
)

type MediaFormat struct {
	ItagNo   int    `json:"itag"`
	MimeType string `json:"mimeType"`

	Width  int `json:"width"`
	Height int `json:"height"`
	FPS    int `json:"fps"`

	AudioSampleRate string `json:"audioSampleRate"`
	AudioChannels   int    `json:"audioChannels"`

	ContentLength int64 `json:"contentLength,string"`
}

type StreamWithCodec struct {
	Stream io.ReadCloser
	Codec  string
}

type Downloader interface {
	VideoFormats(context.Context, string) ([]MediaFormat, error)
	AudioFormats(context.Context, string) ([]MediaFormat, error)

	GetStream(context.Context, string, MediaFormat) (StreamWithCodec, error)
}
