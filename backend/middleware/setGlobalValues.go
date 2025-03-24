package middleware

import (
	"alexlupatsiy.com/personal-website/backend/helpers"
	"github.com/gin-gonic/gin"
)

func SetGlobalValues() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		helpers.SetKV(ctx, "fullPath", ctx.FullPath())
		ctx.Next()
	}
}
