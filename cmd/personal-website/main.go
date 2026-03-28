package main

import (
	"github.com/alexlup06/personal-website/internal/config"
	httpserver "github.com/alexlup06/personal-website/internal/http"
	"github.com/alexlup06/personal-website/internal/http/kit/render"
	"github.com/alexlup06/personal-website/internal/http/middleware"
	"github.com/alexlup06/personal-website/internal/logging"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Version = "dev"

func main() {
	// Binary self-check for Docker HEALTHCHECK
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		os.Exit(0)
	}

	// regular server
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logging.New(cfg.Logging.Level)
	if err != nil {
		logger.Error("invalid token key configuration", "err", err)
		os.Exit(1)
	}
	logger.Info("starting personal-website")

	logging := middleware.RequestLogger(logger)
	setGlobalValues := middleware.SetGlobalValues()
	mw := middleware.Middlewares{
		Logging:         logging,
		SetGlobalValues: setGlobalValues,
	}

	assets, err := render.LoadAssetsManifest("./internal/http/static/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	renderer := render.New(assets)

	server := httpserver.NewServer(httpserver.ServerConfig{
		Version: Version,
		Addr:    cfg.Values.HttpAddr,
		Dev:     cfg.Values.AppEnv == "dev",
		Logger:  logger,
		Render:  renderer,
	}, mw)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("http server listening", "addr", cfg.Values.HttpAddr)
		if err := server.Start(); err != nil {
			logger.Error("http server stopped unexpectedly", "err", err)
			stop()
		}
	}()

	<-ctx.Done()

	logger.Info("shutting down personal-website")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "err", err)
	}

	logger.Info("personal-website stopped")
}
