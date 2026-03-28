package handlers

import (
	"net/http"

	"github.com/alexlup06/personal-website/internal/http/templates/homepage"
	"github.com/alexlup06/personal-website/internal/http/templates/portfolio"
	"github.com/go-chi/chi/v5"
)

func (h *UIHandler) LandingPage(w http.ResponseWriter, r *http.Request) {
	_ = h.Render(
		w,
		r,
		http.StatusOK,
		homepage.Homepage(),
	)
}

func (h *UIHandler) PortfolioPage(w http.ResponseWriter, r *http.Request) {
	_ = h.Render(
		w,
		r,
		http.StatusOK,
		portfolio.Portfolio(),
	)
}

func (h *UIHandler) PorjectPage(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")

	switch project {
	case "authara":
		_ = h.Render(
			w,
			r,
			http.StatusOK,
			portfolio.Authara(),
		)
	}
}
