package domain

type DownloadRequest struct {
	URL string `json:"url"`

	WithVideo bool `json:"with-video"`
	VideoItag int  `json:"video-itag"`

	WithAudio bool `json:"with-audio"`
	AudioItag int  `json:"audio-itag"`
}
