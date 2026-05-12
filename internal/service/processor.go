package service

import (
	"context"
	"sync/atomic"

	"github.com/fwhyjke/youtube/internal/domain"
	"golang.org/x/sync/errgroup"
)

type Task struct {
	id int64

	vStream StreamWithCodec
	aStream StreamWithCodec
}

type Processor struct {
	download Downloader
	convert  Converter

	taskCount atomic.Int64
	tasks     chan Task
}

func NewProcessor(downloader Downloader, converter Converter) *Processor {
	return &Processor{
		download: downloader,
		convert:  converter,
		tasks:    make(chan Task),
	}
}

func (p *Processor) GetVideoFormats(ctx context.Context, url string) ([]domain.MediaFormat, error) {
	return p.download.VideoFormats(ctx, url)
}

func (p *Processor) GetAudioFormats(ctx context.Context, url string) ([]domain.MediaFormat, error) {
	return p.download.AudioFormats(ctx, url)
}

func (p *Processor) CreateTask(ctx context.Context, req domain.DownloadRequest) (int64, error) {
	g, ctx := errgroup.WithContext(ctx)

	t := Task{}

	if req.WithAudio {
		g.Go(func() error {
			var err error
			t.aStream, err = p.download.GetStream(ctx, req.URL, req.AudioItag)
			return err
		})

	}

	if req.WithVideo {
		g.Go(func() error {
			var err error
			t.aStream, err = p.download.GetStream(ctx, req.URL, req.VideoItag)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}

	t.id = p.taskCount.Add(1)

	p.tasks <- t

	return t.id, nil
}
