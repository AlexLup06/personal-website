package backend

import (
	"os"

	"alexlupatsiy.com/personal-website/backend/middleware"
	"alexlupatsiy.com/personal-website/frontend/src/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
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

	r.Use(middleware.CheckHTMXRequest())

	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.Layout())
	})

	return r
}
