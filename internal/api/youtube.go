package api

import (
	"context"

	"github.com/fwhyjke/youtube/internal/domain"
)

type YoutubeService interface {
	getVideoFormats(context.Context, string) ([]domain.MediaFormat, error)
	getAudioFormats(context.Context, string) ([]domain.MediaFormat, error)
}

type YoutubeHTTPHandler struct {
	ytService YoutubeService
}

func NewYoutubeHTTPHandler(ytService YoutubeService) *YoutubeHTTPHandler {
	return &YoutubeHTTPHandler{
		ytService: ytService,
	}
}
