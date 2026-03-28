package httpctx

import "context"

type htmxKeyType struct{}
type fullPathKeyType struct{}
type blogKeyType struct{}

var (
	htmxKey     = htmxKeyType{}
	fullPathKey = fullPathKeyType{}
	blogKey     = blogKeyType{}
)

func WithHTMX(ctx context.Context) context.Context {
	return context.WithValue(ctx, htmxKey, true)
}

func IsHTMX(ctx context.Context) bool {
	v, ok := ctx.Value(htmxKey).(bool)
	return ok && v
}

func WithFullPath(ctx context.Context, fullPath string) context.Context {
	return context.WithValue(ctx, fullPathKey, fullPath)
}

func FullPath(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(fullPathKey).(string)
	return v, ok
}

func WithBlog(ctx context.Context) context.Context {
	return context.WithValue(ctx, blogKey, true)
}

func IsBlog(ctx context.Context) bool {
	v, ok := ctx.Value(htmxKey).(bool)
	return ok && v
}
