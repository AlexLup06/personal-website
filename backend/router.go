package backend

import (
	"context"
	"fmt"

	"alexlupatsiy.com/personal-website/backend/config"
	"alexlupatsiy.com/personal-website/backend/handlers"
	"alexlupatsiy.com/personal-website/backend/middleware"
	"alexlupatsiy.com/personal-website/backend/services"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-envconfig"
)

func Router() error {
	cfg := config.Config{}
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return fmt.Errorf("can't inject env variables: %w", err)
	}

	if !cfg.DevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// services
	blogService := services.NewBlogService()

	// handlers
	router := gin.Default()
	staticHandler := handlers.NewStaticHandler(router)
	homeHandler := handlers.NewHomeHandler(router, blogService)
	portfolioHandler := handlers.NewPortfolioHandler(router)
	blogHandler := handlers.NewBlogHandler(router, blogService)

	// static
	staticHandler.Routes(cfg.DevMode)

	// middleware
	gzipMiddleware := gzip.Gzip(gzip.DefaultCompression)
	checkHTMXMiddleware := middleware.CheckHTMXRequest()
	globalValuesMiddleware := middleware.SetGlobalValues()

	router.Use(gzipMiddleware)
	router.Use(checkHTMXMiddleware)
	router.Use(globalValuesMiddleware)

	// Routes
	homeHandler.Routes()
	portfolioHandler.Routes()
	blogHandler.Routes()

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
