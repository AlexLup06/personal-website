package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	httpmiddleware "alexlupatsiy.com/personal-website/internal/http/middleware"
	"alexlupatsiy.com/personal-website/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Addr        string
	Dev         bool
	Logger      *slog.Logger
	BlogService *services.BlogService
}

func NewServer(cfg Config) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	checkHTMXMiddleware := httpmiddleware.CheckHTMXRequest()
	globalValuesMiddleware := httpmiddleware.SetGlobalValues()

	r.Use(checkHTMXMiddleware)
	r.Use(globalValuesMiddleware)

	r.Use(httpmiddleware.RequestLogger(cfg.Logger))

	registerRoutes(r, cfg)

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{httpServer: srv}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
