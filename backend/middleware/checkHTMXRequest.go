package middleware

import (
	"alexlupatsiy.com/personal-website/backend/helpers"
	"github.com/gin-gonic/gin"
)

func CheckHTMXRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("HX-Request") == "true" {
			helpers.SetKV(ctx, "HX-Request", "true")
		}

		if ctx.GetHeader("X-Blog") == "true" {
			helpers.SetKV(ctx, "X-Blog", "true")
		}
		ctx.Next()
	}
}
