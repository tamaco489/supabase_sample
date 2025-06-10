package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/tamaco489/supabase_sample/api/shop/intrnal/handler"
)

func main() {
	ctx := context.Background()

	http.HandleFunc("/api/healthcheck", handler.HealthCheckHandler)

	slog.InfoContext(ctx, "Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
