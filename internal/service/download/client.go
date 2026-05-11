package service

import (
	"context"

	"github.com/fwhyjke/youtube/internal/service"
	"github.com/kkdai/youtube/v2"
)

type Client struct {
	client *youtube.Client
}

func NewClient() *Client {
	return &Client{
		client: &youtube.Client{},
	}
}

func (c Client) AudioFormats(ctx context.Context, url string) ([]service.MediaFormat, error) {
	aFormats := []service.MediaFormat{}

	meta, err := c.client.GetVideoContext(ctx, url)
	if err != nil {
		return aFormats, err
	}

	for _, format := range meta.Formats.Type("audio/mp4") {
		aFormats = append(aFormats, service.MediaFormat{
			ItagNo:   format.ItagNo,
			MimeType: format.MimeType,

			AudioSampleRate: format.AudioSampleRate,
			AudioChannels:   format.AudioChannels,

			ContentLength: format.ContentLength,
		})
	}

	return aFormats, err
}

func (c Client) VideoFormats(ctx context.Context, url string) ([]service.MediaFormat, error) {
	vFormats := []service.MediaFormat{}

	meta, err := c.client.GetVideoContext(ctx, url)
	if err != nil {
		return vFormats, err
	}

	for _, format := range meta.Formats.Type("video/mp4") {
		vFormats = append(vFormats, service.MediaFormat{
			ItagNo:   format.ItagNo,
			MimeType: format.MimeType,

			Width:  format.Width,
			Height: format.Height,
			FPS:    format.FPS,

			ContentLength: format.ContentLength,
		})
	}

	return vFormats, err
}

func (c Client) GetStream(ctx context.Context, url string, formatItag int) (service.StreamWithCodec, error) {
	meta, err := c.client.GetVideoContext(ctx, url)
	if err != nil {
		return service.StreamWithCodec{}, err
	}

	f := &meta.Formats.Itag(formatItag)[0]
	stream, _, err := c.client.GetStream(meta, f)
	if err != nil {
		return service.StreamWithCodec{}, err
	}

	return service.StreamWithCodec{
		Stream: stream,
		Codec:  f.MimeType,
	}, nil
}
