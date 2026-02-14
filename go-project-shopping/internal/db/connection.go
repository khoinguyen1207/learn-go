package db

import (
	"context"
	"fmt"
	"log"
	"project-shopping/internal/config"
	"project-shopping/internal/db/sqlc"
	"project-shopping/internal/utils"
	"project-shopping/pkg/pgx"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

var DB sqlc.Querier

func InitDB(cfg *config.Config) error {
	dns := cfg.Db.DatabaseUrl

	poolConfig, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return fmt.Errorf("Unable to parse database configuration: %v", err)
	}

	sqlLogger := utils.NewLoggerWithPath("internal/logs/sql.log", "info")

	// Configure pgx tracer with zerolog
	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: &pgx.PgxZeroLogTracer{
			Logger:         *sqlLogger,
			SlowQueryLimit: 500 * time.Millisecond,
		},
		LogLevel: tracelog.LogLevelDebug,
	}

	poolConfig.MaxConns = 30
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 30 * time.Minute  // 30 minutes
	poolConfig.MaxConnIdleTime = 5 * time.Minute   // 5 minutes
	poolConfig.HealthCheckPeriod = 1 * time.Minute // 1 minutes

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbpool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("Unable to create database connection pool: %v", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return fmt.Errorf("Could not connect to the database: %v", err)
	}

	DB = sqlc.New(dbpool)

	log.Println("Connected to the database successfully!")

	return nil
}
