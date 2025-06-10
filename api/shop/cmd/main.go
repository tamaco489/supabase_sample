package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"OK"}`))
}

func main() {
	ctx := context.Background()

	http.HandleFunc("/api/healthcheck", healthCheckHandler)

	slog.InfoContext(ctx, "Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
