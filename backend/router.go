package backend

import (
	"alexlupatsiy.com/personal-website/frontend/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func Router() *gin.Engine {
	r := gin.Default()

	r.Static("/js", "./frontend/js")
	r.Static("/css", "./frontend/css")

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
