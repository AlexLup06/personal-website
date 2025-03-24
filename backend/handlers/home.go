package handlers

import (
	"alexlupatsiy.com/personal-website/backend/helpers"
	"alexlupatsiy.com/personal-website/backend/services"
	"alexlupatsiy.com/personal-website/frontend/src/views/homepage"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	router      *gin.Engine
	blogService *services.BlogService
}

func NewHomeHandler(
	router *gin.Engine,
	blogService *services.BlogService,
) *HomeHandler {
	return &HomeHandler{
		router:      router,
		blogService: blogService,
	}
}

func (h *HomeHandler) Routes() {
	homeRouter := h.router.Group("")

	homeRouter.GET("", h.landingPage)
}

func (p *HomeHandler) landingPage(ctx *gin.Context) {
	showcasedBlogs := p.blogService.GetMostRecentBlogs()
	helpers.Render(ctx, 200, homepage.Homepage(showcasedBlogs))
}
