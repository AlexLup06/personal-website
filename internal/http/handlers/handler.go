package handlers

import (
	"github.com/alexlup06/personal-website/internal/http/kit/render"
)

type UIHandlerConfig struct {
	Render render.Renderer
}

type UIHandler struct {
	Render render.Renderer
}

func NewUIHandler(cfg UIHandlerConfig) *UIHandler {
	return &UIHandler{
		Render: cfg.Render,
	}
}
