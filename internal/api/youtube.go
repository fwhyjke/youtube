package api

import (
	"context"
	"net/http"

	"github.com/fwhyjke/youtube/internal/api/server"
	"github.com/fwhyjke/youtube/internal/domain"
)

type YoutubeService interface {
	GetVideoFormats(context.Context, string) ([]domain.MediaFormat, error)
	GetAudioFormats(context.Context, string) ([]domain.MediaFormat, error)
}

type YoutubeHTTPHandler struct {
	ytService YoutubeService
}

func NewYoutubeHTTPHandler(ytService YoutubeService) *YoutubeHTTPHandler {
	return &YoutubeHTTPHandler{
		ytService: ytService,
	}
}

func (h *YoutubeHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "get-video-formats",
			Handler: h.GetVideoFormats,
		},
		{
			Method:  http.MethodGet,
			Path:    "get-audio-formats",
			Handler: h.GetAudioFormats,
		},
	}
}
