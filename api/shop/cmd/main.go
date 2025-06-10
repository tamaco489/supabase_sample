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

	http.HandleFunc("/shop/v1/healthcheck", handler.HealthCheckHandler)
	http.HandleFunc("/shop/v1/users/me", handler.GetMeMock)

	slog.InfoContext(ctx, "Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
