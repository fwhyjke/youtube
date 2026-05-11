package download

import (
	"context"
	"io"

	"github.com/fwhyjke/youtube/internal/api"
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

func (c Client) AudioFormats(url string) ([]api.MediaFormat, error) {
	aFormats := []api.MediaFormat{}

	meta, err := c.client.GetVideoContext(context.TODO(), url)
	if err != nil {
		return aFormats, err
	}

	for _, format := range meta.Formats.Type("audio/mp4") {
		aFormats = append(aFormats, api.MediaFormat{
			ItagNo:   format.ItagNo,
			MimeType: format.MimeType,

			AudioSampleRate: format.AudioSampleRate,
			AudioChannels:   format.AudioChannels,

			ContentLength: format.ContentLength,
		})
	}

	return aFormats, err
}

func (c Client) VideoFormats(url string) ([]api.MediaFormat, error) {
	vFormats := []api.MediaFormat{}

	meta, err := c.client.GetVideoContext(context.TODO(), url)
	if err != nil {
		return vFormats, err
	}

	for _, format := range meta.Formats.Type("video/mp4") {
		vFormats = append(vFormats, api.MediaFormat{
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

func (c Client) GetStream(url string, format api.MediaFormat) (io.ReadCloser, string, error) {
	meta, err := c.client.GetVideoContext(context.TODO(), url)
	if err != nil {
		return nil, "", err
	}

	stream, _, err := c.client.GetStream(meta, &meta.Formats.Itag(format.ItagNo)[0])
	if err != nil {
		return nil, "", err
	}

	return stream, format.MimeType, nil
}
