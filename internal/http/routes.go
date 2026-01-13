package http

import (
	"alexlupatsiy.com/personal-website/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r chi.Router, cfg Config) {
	homeHandler := handlers.NewHomeHandler(cfg.BlogService)
	portfolioHandler := handlers.NewPortfolioHandler()
	blogHandler := handlers.NewBlogHandler(cfg.BlogService)

	homeHandler.Routes(r)
	portfolioHandler.Routes(r)
	blogHandler.Routes(r)

	handlers.RegisterStatic(r, handlers.StaticConfig{
		Dev: cfg.Dev,
	})

}
