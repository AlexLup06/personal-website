package backend

import (
	"alexlupatsiy.com/personal-website/backend/helpers"
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
	r := gin.Default()

	envVars := helpers.LoadEnv()
	mode := envVars["MODE"]
	isProductionMode := mode == "production"

	static := r.Group("/", middleware.ServeGzippedFiles(isProductionMode))
	{
		static.GET("/js/*filepath", middleware.ServeStaticFiles("./frontend/src/js"))
		static.GET("/css/*filepath", middleware.ServeStaticFiles("./frontend/src/css"))
		static.GET("/public/*filepath", middleware.ServeStaticFiles("./frontend/public"))
	}

	r.Use(middleware.CheckHTMXRequest())

	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.Home())
	})

	return r
}
