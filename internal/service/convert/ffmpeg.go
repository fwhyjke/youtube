package convert

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fwhyjke/youtube/internal/service"
	"golang.org/x/sync/errgroup"
)

type FFmpegStore struct {
	tmpDir    string
	resultDir string
}

func NewFFmpegStore(temp, out string) *FFmpegStore {
	return &FFmpegStore{
		tmpDir:    temp,
		resultDir: out,
	}
}

func (f *FFmpegStore) pullStream(ctx context.Context, stream io.ReadCloser) (string, error) {
	defer stream.Close()

	path := fmt.Sprintf("%s/%d", f.tmpDir, time.Now().UnixNano())

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}

	return path, err
}

func (f *FFmpegStore) BuildOnlyA(ctx context.Context, stream service.StreamWithCodec) (string, error) {
	iPath, err := f.pullStream(ctx, stream.Stream)
	if err != nil {
		return "", err
	}
	defer os.Remove(iPath)

	outPath := fmt.Sprintf("%s/%d.mp3", f.resultDir, time.Now().UnixNano())

	cmd := exec.CommandContext(ctx, "ffmpeg", "-y",
		"-i", iPath,
		"-vn",
		"-c:a", "libmp3lame",
		"-q:a", "2",
		outPath,
	)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outPath, err
}

func (f *FFmpegStore) BuildOnlyV(ctx context.Context, stream service.StreamWithCodec) (string, error) {
	iPath, err := f.pullStream(ctx, stream.Stream)
	if err != nil {
		return "", err
	}
	defer os.Remove(iPath)

	outPath := fmt.Sprintf("%s/%d.mp4", f.resultDir, time.Now().UnixNano())

	var cmd *exec.Cmd
	if strings.Contains(stream.Codec, "avc1") {
		cmd = exec.CommandContext(ctx, "ffmpeg", "-y",
			"-i", iPath,
			"-c", "copy",
			"-map", "0:v:0",
			outPath,
		)
	} else {
		cmd = exec.CommandContext(ctx, "ffmpeg", "-y",
			"-i", iPath,
			"-c:v", "libx264",
			"-crf", "18",
			"-preset", "fast",
			"-map", "0:v:0",
			outPath,
		)
	}

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outPath, err
}

func (f *FFmpegStore) BuildVandA(ctx context.Context, vStream service.StreamWithCodec, aStream service.StreamWithCodec) (string, error) {
	g, ctx := errgroup.WithContext(ctx)

	var (
		vPath string
		aPath string
	)

	g.Go(func() error {
		var err error

		vPath, err = f.pullStream(ctx, vStream.Stream)

		return err
	})

	g.Go(func() error {
		var err error

		aPath, err = f.pullStream(ctx, aStream.Stream)

		return err
	})

	if err := g.Wait(); err != nil {
		return "", err
	}

	defer os.Remove(vPath)
	defer os.Remove(aPath)

	outPath := fmt.Sprintf("%s/%d.mp4", f.resultDir, time.Now().UnixNano())

	var cmd *exec.Cmd
	if strings.Contains(vStream.Codec, "avc1") {
		cmd = exec.CommandContext(ctx, "ffmpeg", "-y",
			"-i", vPath,
			"-i", aPath,
			"-c", "copy",
			"-map", "0:v:0",
			"-map", "1:a:0",
			"-shortest",
			outPath,
		)
	} else {
		cmd = exec.CommandContext(ctx, "ffmpeg", "-y",
			"-i", vPath,
			"-i", aPath,
			"-c:v", "libx264",
			"-crf", "18",
			"-preset", "fast",
			"-c:a", "copy",
			"-map", "0:v:0",
			"-map", "1:a:0",
			"-shortest",
			outPath,
		)
	}

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return outPath, err
}
