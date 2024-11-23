package middleware

import "github.com/gin-gonic/gin"

func CheckHTMXRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "true" {
			c.Set("HX-Request", true)
		}
		c.Next()
	}
}
