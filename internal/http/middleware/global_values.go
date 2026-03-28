package middleware

import (
	"net/http"

	"github.com/alexlup06/personal-website/internal/http/kit/httpctx"
	"github.com/go-chi/chi/v5"
)

func SetGlobalValues() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			routeCtx := chi.RouteContext(ctx)
			if routeCtx != nil {
				pattern := routeCtx.RoutePattern()
				if pattern != "" {
					ctx = httpctx.WithFullPath(ctx, pattern)
				}
			}

			if r.Header.Get("HX-Request") == "true" {
				ctx = httpctx.WithHTMX(ctx)
			}

			if r.Header.Get("X-Blog") == "true" {
				ctx = httpctx.WithBlog(ctx)
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
