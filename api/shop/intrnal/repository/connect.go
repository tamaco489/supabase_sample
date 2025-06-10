package repository

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once   sync.Once
)

const dsn = ""

func InitDB(ctx context.Context) {
	once.Do(func() {
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Failed to parse DB config: %v", err)
		}

		config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

		dbPool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})
}

func GetPool() *pgxpool.Pool {
	return dbPool
}
