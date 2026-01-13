package middleware

import (
	"net/http"

	"alexlupatsiy.com/personal-website/internal/helpers"
	"github.com/go-chi/chi/v5"
)

func SetGlobalValues() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// chi stores route context here
			routeCtx := chi.RouteContext(r.Context())

			if routeCtx != nil {
				pattern := routeCtx.RoutePattern()
				if pattern != "" {
					ctx := helpers.WithKV(
						r.Context(),
						helpers.CtxFullPath,
						pattern,
					)
					r = r.WithContext(ctx)
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
