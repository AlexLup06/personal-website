package backend

import (
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

	r.Static("/js", "./frontend/src/js")
	r.Static("/css", "./frontend/src/css")
	r.Static("/public", "./frontend/public")

	r.Use(middleware.CheckHTMXRequest())

	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.Layout())
	})

	r.GET("/test1", func(c *gin.Context) {
		render(c, 200, views.Test2())
	})

	r.GET("/test2", func(c *gin.Context) {
		render(c, 200, views.Test1())
	})

	return r
}
