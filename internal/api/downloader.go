package api

import (
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

type Downloader interface {
	VideoFormats(string) ([]MediaFormat, error)
	AudioFormats(string) ([]MediaFormat, error)

	GetStream(string, MediaFormat) (io.ReadCloser, string, error)
}
