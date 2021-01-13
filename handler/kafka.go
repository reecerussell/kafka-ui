package handler

import (
	"encoding/json"
	"net/http"

	"github.com/reecerussell/kafka-ui/config"
	"github.com/reecerussell/kafka-ui/model"
)

// GetKafka returns a http.Handler which returns Kafka config.
func GetKafkaSettings(cnf *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cnf.Kafka)
	})
}

//
func SetKafkaSettings(cnf *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data model.Kafka
		_ = json.NewDecoder(r.Body).Decode(&data)
		defer r.Body.Close()

		err := data.Validate()
		if err != nil {
			resp := model.Error{Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
			return
		}

		cnf.Kafka = &data
		err = cnf.Save()
		if err != nil {
			resp := model.Error{Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cnf.Kafka)
	})
}
