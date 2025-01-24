package backend

import (
	"os"

	"alexlupatsiy.com/personal-website/backend/middleware"
	"alexlupatsiy.com/personal-website/frontend/src/views"
	"alexlupatsiy.com/personal-website/frontend/src/views/blog"
	"alexlupatsiy.com/personal-website/frontend/src/views/homepage"
	"alexlupatsiy.com/personal-website/frontend/src/views/portfolio"
	"github.com/a-h/templ"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c, c.Writer)
}

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

	r := gin.Default()

	static := r.Group("/", middleware.ServeGzippedFiles(isProductionMode))
	{
		static.GET("/js/*filepath", middleware.ServeStaticFiles("./frontend/src/js"))
		static.GET("/css/*filepath", middleware.ServeStaticFiles("./frontend/src/css"))
		static.GET("/public/*filepath", middleware.ServeStaticFiles(staticBasePath))
	}

	r.Use(middleware.CheckHTMXRequest(), middleware.SetGlobalValues())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", func(c *gin.Context) {
		render(c, 200, homepage.Homepage())
	})

	r.GET("/portfolio", func(c *gin.Context) {
		render(c, 200, portfolio.Portfolio())
	})

	r.GET("/blog", func(c *gin.Context) {
		render(c, 200, blog.Blog())
	})

	if !isProductionMode {
		r.GET("/test", func(c *gin.Context) {
			render(c, 200, views.Test())
		})
	}
	return r
}
