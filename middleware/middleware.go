package middleware

import "net/http"

// Constructor is a function used to create middleware.
type Constructor func(h http.Handler) http.Handler

// Pipe combines all middleware and the given handler(h).
func Pipe(h http.Handler, middleware ...Constructor) http.Handler {
	base := h

	for i := 0; i < len(middleware); i++ {
		base = middleware[i](base)
	}

	return base
}
