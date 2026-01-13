package helpers

import "context"

// private key type (still prevents collisions)
type ctxKey string

// ------------------------------------------------------------
// EXPORTED keys (capitalized)
// ------------------------------------------------------------

const (
	CtxHXRequest ctxKey = "hx-request"
	CtxBlog      ctxKey = "x-blog"
	CtxFilters   ctxKey = "filters"
	CtxFullPath  ctxKey = "full-path"
)

// ------------------------------------------------------------
// Generic helpers
// ------------------------------------------------------------

func WithKV(ctx context.Context, key ctxKey, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetKV[T any](ctx context.Context, key ctxKey) (T, bool) {
	v, ok := ctx.Value(key).(T)
	return v, ok
}

// ------------------------------------------------------------
// Filters helpers (typed)
// ------------------------------------------------------------

func WithFilters(ctx context.Context, filters []string) context.Context {
	return context.WithValue(ctx, CtxFilters, filters)
}

func Filters(ctx context.Context) ([]string, bool) {
	v, ok := ctx.Value(CtxFilters).([]string)
	return v, ok
}

// ------------------------------------------------------------
// Request-type helpers
// ------------------------------------------------------------

func WithHTMX(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxHXRequest, true)
}

func WithBlog(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxBlog, true)
}

func IsHTMX(ctx context.Context) bool {
	v, ok := GetKV[bool](ctx, CtxHXRequest)
	return ok && v
}

func IsBlog(ctx context.Context) bool {
	v, ok := GetKV[bool](ctx, CtxBlog)
	return ok && v
}
