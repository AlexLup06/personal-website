package handlers

import (
	"fmt"
	"net/http"

	"alexlupatsiy.com/personal-website/internal/helpers"
	"alexlupatsiy.com/personal-website/internal/http/templates/homepage"
	"alexlupatsiy.com/personal-website/internal/services"
	"github.com/go-chi/chi/v5"
)

type HomeHandler struct {
	blogService *services.BlogService
}

func NewHomeHandler(blogService *services.BlogService) *HomeHandler {
	return &HomeHandler{
		blogService: blogService,
	}
}

// Routes registers routes on the provided chi router
func (h *HomeHandler) Routes(r chi.Router) {
	r.Get("/", h.landingPage)
}

func (h *HomeHandler) landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(helpers.IsHTMX(r.Context()))
	_ = helpers.Render(
		w,
		r,
		http.StatusOK,
		homepage.Homepage(),
	)
}
