package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"alexlupatsiy.com/personal-website/internal/config"
	httpserver "alexlupatsiy.com/personal-website/internal/http"
	"alexlupatsiy.com/personal-website/internal/logging"
	"alexlupatsiy.com/personal-website/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logging.New(cfg.Logging.Level)
	logger.Info("starting authgate")

	blogService := services.NewBlogService()

	server := httpserver.NewServer(httpserver.Config{
		Addr:        cfg.HTTP.Addr,
		Dev:         cfg.Dev,
		BlogService: blogService,
		Logger:      logger,
	})

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("http server listening", "addr", cfg.HTTP.Addr)
		if err := server.Start(); err != nil {
			logger.Error("http server stopped unexpectedly", "err", err)
			stop()
		}
	}()

	<-ctx.Done()

	logger.Info("shutting down authgate")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "err", err)
	}

	logger.Info("authgate stopped")
}
