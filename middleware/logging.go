package middleware

import (
	"fmt"
	"net/http"

	"github.com/reecerussell/kafka-ui/logging"
)

// NewLoggingMiddleware returns middleware that logs incoming HTTP requests.
func NewLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery)
		logging.Info("%s: %s", r.Method, path)

		next.ServeHTTP(w, r)
	})
}
