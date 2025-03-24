package helpers

import (
	"context"

	"github.com/gin-gonic/gin"
)

const FILTER_KEY = "filters"

func SetKV(ctx *gin.Context, key, value string) {
	newCtx := context.WithValue(ctx.Request.Context(), key, value)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

func SetContextWithFilters(ctx *gin.Context, value []string) {
	newCtx := context.WithValue(ctx.Request.Context(), FILTER_KEY, value)
	ctx.Request = ctx.Request.WithContext(newCtx)
}
