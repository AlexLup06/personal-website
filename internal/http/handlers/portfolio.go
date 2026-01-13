package handlers

import (
	"net/http"

	"alexlupatsiy.com/personal-website/internal/helpers"
	"alexlupatsiy.com/personal-website/internal/http/templates/portfolio"
	"github.com/go-chi/chi/v5"
)

type PortfolioHandler struct{}

func NewPortfolioHandler() *PortfolioHandler {
	return &PortfolioHandler{}
}

// Routes registers portfolio routes on the provided router
func (p *PortfolioHandler) Routes(r chi.Router) {
	r.Route("/portfolio", func(r chi.Router) {
		r.Get("/", p.landingPage)
	})
}

func (p *PortfolioHandler) landingPage(w http.ResponseWriter, r *http.Request) {
	_ = helpers.Render(
		w,
		r,
		http.StatusOK,
		portfolio.Portfolio(),
	)
}
