package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/alexlup06/personal-website/internal/http/kit/render"
	httpmiddleware "github.com/alexlup06/personal-website/internal/http/middleware"
)

type ServerConfig struct {
	Version string
	Addr    string
	Dev     bool
	Logger  *slog.Logger
	Render  render.Renderer
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg ServerConfig, mw httpmiddleware.Middlewares) *Server {
	handler := NewRouter(cfg, mw)

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
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
