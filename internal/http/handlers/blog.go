package handlers

import (
	"net/http"
	"slices"
	"strings"

	"alexlupatsiy.com/personal-website/internal/helpers"
	"alexlupatsiy.com/personal-website/internal/http/templates/blog"
	"alexlupatsiy.com/personal-website/internal/services"
	"github.com/go-chi/chi/v5"
)

type BlogHandler struct {
	blogService *services.BlogService
}

func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

// Routes registers blog routes on the provided chi router
func (b *BlogHandler) Routes(r chi.Router) {
	r.Route("/blog", func(r chi.Router) {
		r.Get("/", b.blogOverviewGET)
		r.Get("/{slug}", b.blogGET)
	})
}

func (b *BlogHandler) blogOverviewGET(w http.ResponseWriter, r *http.Request) {
	// query params: ?filter=a&filter=b
	selectedCategories := r.URL.Query()["filter"]

	if slices.Contains(selectedCategories, "all") {
		selectedCategories = []string{}
	}

	blogMetadata := b.blogService.GetFilteredMetadata(selectedCategories)

	// attach filters to request context
	ctx := helpers.WithFilters(r.Context(), selectedCategories)
	r = r.WithContext(ctx)

	// set HX-Replace-Url
	setBlogURL(w, selectedCategories)

	// HTMX partial update case
	isBlog := helpers.IsBlog(r.Context())
	isHX := helpers.IsHTMX(r.Context())

	if isBlog && isHX {
		w.Header().Set("HX-Reswap", "none")
		_ = helpers.Render(
			w,
			r,
			http.StatusOK,
			blog.FilterAndCards(blogMetadata),
		)
		return
	}

	// full page render
	_ = helpers.Render(
		w,
		r,
		http.StatusOK,
		blog.BlogLandingPage(blogMetadata),
	)
}

func setBlogURL(w http.ResponseWriter, selectedCategories []string) {
	if len(selectedCategories) == 0 {
		w.Header().Set("HX-Replace-Url", "/blog")
		return
	}

	var sb strings.Builder
	sb.WriteString("/blog?")

	for i, c := range selectedCategories {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString("filter=")
		sb.WriteString(c)
	}

	w.Header().Set("HX-Replace-Url", sb.String())
}

func (b *BlogHandler) blogGET(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	blogMetadata, exists := b.blogService.GetBlogMetadata(slug)
	if !exists {
		http.NotFound(w, r)
		return
	}

	content, err := b.blogService.GetBlogContent(slug)
	if err != nil {
		http.Error(w, "Error reading blog content", http.StatusInternalServerError)
		return
	}

	_ = helpers.Render(
		w,
		r,
		http.StatusOK,
		blog.Blog(
			blogMetadata.Title,
			blogMetadata.Date,
			blogMetadata.Intro,
			content,
		),
	)
}
