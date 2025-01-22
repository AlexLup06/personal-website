package middleware

import "github.com/gin-gonic/gin"

func SetGlobalValues() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("fullPath", c.FullPath())
		c.Next()
	}
}
