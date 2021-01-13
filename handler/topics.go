package handler

import (
	"encoding/json"
	"net/http"

	"github.com/reecerussell/kafka-ui/config"
)

// GetTopics returns a http.Handler which returns Topics
// from the given config.
func GetTopics(cnf *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cnf.Topics)
	})
}