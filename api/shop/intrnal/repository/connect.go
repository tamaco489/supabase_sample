package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DBConfig is a struct that holds database configuration.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

// GetDBConfig returns a DBConfig struct with the database configuration.
func GetDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

// getDSN returns a DSN string for the database.
func getDSN(cfg DBConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

var (
	// dbPool is the database connection pool.
	dbPool *pgxpool.Pool

	// once is used to ensure that the database connection pool is initialized only once.
	once sync.Once
)

// InitDB initializes the database connection pool.
func InitDB(ctx context.Context) {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(getDSN(GetDBConfig()))
		if err != nil {
			log.Fatalf("Failed to parse DB config: %v", err)
		}

		// NOTE: Changed to SimpleProtocol as it defaults to cache_statement
		cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

		// Set maximum number of DB connections to 2
		cfg.MaxConns = 2

		// Set maximum connection lifetime to 1 minute
		cfg.MaxConnLifetime = 1 * time.Minute

		// Create a new database connection pool
		dbPool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})
}

// GetPool returns the database connection pool.
func GetPool() *pgxpool.Pool {
	return dbPool
}
