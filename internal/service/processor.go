package service

import (
	"context"

	"github.com/fwhyjke/youtube/internal/domain"
)

type Processor struct {
	download Downloader
	convert  Converter
}

func NewProcessor(downloader Downloader, converter Converter) Processor {
	return Processor{
		download: downloader,
		convert:  converter,
	}
}

func (p Processor) GetVideoFormats(ctx context.Context, url string) ([]domain.MediaFormat, error) {
	return p.download.VideoFormats(ctx, url)
}

func (p Processor) GetAudioFormats(ctx context.Context, url string) ([]domain.MediaFormat, error) {
	return p.download.AudioFormats(ctx, url)
}
