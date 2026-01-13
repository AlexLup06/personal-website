package middleware

import (
	"net/http"

	"alexlupatsiy.com/personal-website/internal/helpers"
)

func CheckHTMXRequest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			if r.Header.Get("HX-Request") == "true" {
				ctx = helpers.WithHTMX(ctx)
			}

			if r.Header.Get("X-Blog") == "true" {
				ctx = helpers.WithBlog(ctx)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
