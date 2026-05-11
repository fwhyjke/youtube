package service

import (
	"context"
	"io"

	"github.com/fwhyjke/youtube/internal/domain"
)

type StreamWithCodec struct {
	Stream io.ReadCloser
	Codec  string
}

type Downloader interface {
	VideoFormats(context.Context, string) ([]domain.MediaFormat, error)
	AudioFormats(context.Context, string) ([]domain.MediaFormat, error)

	GetStream(context.Context, string, int) (StreamWithCodec, error)
}
