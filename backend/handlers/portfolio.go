package handlers

import (
	"alexlupatsiy.com/personal-website/backend/helpers"
	"alexlupatsiy.com/personal-website/frontend/src/views/portfolio"
	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	router *gin.Engine
}

func NewPortfolioHandler(
	router *gin.Engine,
) *PortfolioHandler {
	return &PortfolioHandler{
		router: router,
	}
}

func (p *PortfolioHandler) Routes() {
	portfolioRouter := p.router.Group("/portfolio")

	portfolioRouter.GET("", p.landingPage)
}

func (p *PortfolioHandler) landingPage(ctx *gin.Context) {
	helpers.Render(ctx, 200, portfolio.Portfolio())

}
