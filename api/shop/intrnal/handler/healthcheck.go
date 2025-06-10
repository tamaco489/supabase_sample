package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
