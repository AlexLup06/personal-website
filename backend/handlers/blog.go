package handlers

import (
	"net/http"
	"slices"

	"alexlupatsiy.com/personal-website/backend/helpers"
	"alexlupatsiy.com/personal-website/backend/services"
	"alexlupatsiy.com/personal-website/frontend/src/views/blog"
	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	router      *gin.Engine
	blogService *services.BlogService
}

func NewBlogHandler(
	router *gin.Engine,
	blogService *services.BlogService,
) *BlogHandler {
	return &BlogHandler{
		router:      router,
		blogService: blogService,
	}
}

func (b *BlogHandler) Routes() {
	blogRouter := b.router.Group("/blog")

	blogRouter.GET("", b.blogOverviewGET)
	blogRouter.GET("/:slug", b.blogGET)
}

func (b *BlogHandler) blogOverviewGET(ctx *gin.Context) {
	selectedCategories := ctx.QueryArray("filter")

	if slices.Contains(selectedCategories, "all") {
		selectedCategories = []string{}
	}

	blogMetadata := b.blogService.GetFilteredMetadata(selectedCategories)
	helpers.SetContextWithFilters(ctx, selectedCategories)
	setUrl(ctx, selectedCategories)

	// User applied filter and we only send BlogCards and filter component
	if ctx.Request.Context().Value("X-Blog") == "true" && ctx.Request.Context().Value("HX-Request") == "true" {
		ctx.Header("HX-Reswap", "none")
		helpers.Render(ctx, 200, blog.FilterAndCards(blogMetadata))
		return
	}

	helpers.Render(ctx, 200, blog.BlogLandingPage(blogMetadata))
}

func setUrl(ctx *gin.Context, selectedCategories []string) {
	url := "/blog"
	for i, selectedCategorie := range selectedCategories {
		if i == 0 {
			url += "?"
		}
		url += "filter=" + selectedCategorie
		if i != len(selectedCategories)-1 {
			url += "&"
		}
	}

	ctx.Header("HX-Replace-Url", url)
}

func (b *BlogHandler) blogGET(ctx *gin.Context) {
	slug := ctx.Param("slug")

	blogMetadata, exists := b.blogService.GetBlogMetadata(slug)
	if !exists {
		ctx.String(http.StatusNotFound, "Blog not found")
		return
	}

	content, err := b.blogService.GetBlogContent(slug)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error reading blog content")
		return
	}

	helpers.Render(ctx, 200, blog.Blog(blogMetadata.Title, blogMetadata.Date, blogMetadata.Intro, content))
}
