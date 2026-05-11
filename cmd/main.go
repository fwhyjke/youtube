package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/fwhyjke/youtube/internal/api"
	"github.com/fwhyjke/youtube/internal/api/server"
	"github.com/fwhyjke/youtube/internal/service"
	"github.com/fwhyjke/youtube/internal/service/convert"
	"github.com/fwhyjke/youtube/internal/service/download"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	downloader := download.NewClient()
	conv := convert.NewFFmpegStore("/tmp", "/result")

	ytHandler := api.NewYoutubeHTTPHandler(service.NewProcessor(downloader, conv))
	ytRoutes := ytHandler.Routes()

	router := server.NewAPIRouter()
	router.AddRoutes(ytRoutes...)

	cfg, err := server.NewConfig()
	if err != nil {
		panic(err)
	}

	serv := server.NewHTTPServer(cfg)
	serv.RegisterAPIRouters(router)

	err = serv.Run(ctx)
	if err != nil {
		panic(err)
	}
}
