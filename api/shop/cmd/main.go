package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tamaco489/supabase_sample/api/shop/intrnal/handler"
	"github.com/tamaco489/supabase_sample/api/shop/intrnal/repository"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Initialize the database connection pool
	ctx := context.Background()
	repository.InitDB(ctx)

	// Register HTTP handlers
	http.HandleFunc("/shop/v1/healthcheck", handler.HealthCheckHandler)
	http.HandleFunc("/shop/v1/users/me", handler.GetMe)

	// Start the HTTP server
	env, port, service := os.Getenv("API_ENV"), os.Getenv("API_PORT"), os.Getenv("API_SERVICE_NAME")
	slog.InfoContext(ctx, fmt.Sprintf("[%s] %s starting on %s", env, service, port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
