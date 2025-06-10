package repository

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once   sync.Once
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		User:     "",
		Password: "",
		Host:     "",
		Port:     "",
		DBName:   "",
		SSLMode:  "verify-full",
	}
}

func getDSN(config DBConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
}

func InitDB(ctx context.Context) {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(getDSN(GetDBConfig()))
		if err != nil {
			log.Fatalf("Failed to parse DB config: %v", err)
		}

		// NOTE: デフォルトでは cache_statement になっているので、SimpleProtocol に変更
		cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
		cfg.MaxConns = 2
		cfg.MaxConnLifetime = 1 * time.Minute

		dbPool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})
}

func GetPool() *pgxpool.Pool {
	return dbPool
}
