package backend

import (
	"fmt"
	"os"

	"alexlupatsiy.com/personal-website/backend/handlers"
	"alexlupatsiy.com/personal-website/backend/helpers"
	"alexlupatsiy.com/personal-website/backend/middleware"
	"alexlupatsiy.com/personal-website/frontend/src/views"
	"alexlupatsiy.com/personal-website/frontend/src/views/blog"
	"alexlupatsiy.com/personal-website/frontend/src/views/homepage"
	"alexlupatsiy.com/personal-website/frontend/src/views/portfolio"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	env := os.Getenv("ENV")
	isProductionMode := env == "production"
	if isProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	var staticBasePath string
	if isProductionMode {
		staticBasePath = "/root/public"
	} else {
		staticBasePath = "./frontend/public"
	}

	var blogMetadata map[string]blog.BlogType
	blogMetadata, err := handlers.LoadBlogMetadata()
	if err != nil {
		fmt.Printf("Error reading blog metadata: %v\n", err)
	}

	r := gin.Default()
	static := r.Group("/", middleware.ServeGzippedFiles(isProductionMode))
	{
		static.GET("/js/*filepath", middleware.ServeStaticFiles("./frontend/src/js"))
		static.GET("/css/*filepath", middleware.ServeStaticFiles("./frontend/src/css"))
		static.GET("/public/*filepath", middleware.ServeStaticFiles(staticBasePath))
	}

	r.Use(middleware.CheckHTMXRequest(), middleware.SetGlobalValues())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", func(c *gin.Context) { helpers.Render(c, 200, homepage.Homepage()) })
	r.GET("/portfolio", func(c *gin.Context) { helpers.Render(c, 200, portfolio.Portfolio()) })
	blogRouter := r.Group("/blog")
	{
		blogRouter.GET("", func(c *gin.Context) { helpers.Render(c, 200, blog.BlockLandingPage(blogMetadata)) })
		blogRouter.GET("/:slug", func(ctx *gin.Context) { handlers.BlogHandler(ctx, blogMetadata) })
	}

	if !isProductionMode {
		r.GET("/test", func(c *gin.Context) {
			helpers.Render(c, 200, views.Test())
		})
	}
	return r
}
