package middleware

import "net/http"

type Middlewares struct {
	Logging         func(http.Handler) http.Handler
	SetGlobalValues func(http.Handler) http.Handler
}
