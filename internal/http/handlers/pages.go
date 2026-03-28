package handlers

import (
	"net/http"

	"github.com/alexlup06/personal-website/internal/http/templates/homepage"
	"github.com/alexlup06/personal-website/internal/http/templates/portfolio"
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
