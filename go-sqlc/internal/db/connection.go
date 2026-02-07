package db

import (
	"context"
	"fmt"
	"go-sqlc/internal/configs"
	"go-sqlc/internal/db/sqlc"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *sqlc.Queries

func InitDB(cfg configs.Config) error {
	dns := cfg.Db.ConnectionString

	poolConfig, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return fmt.Errorf("Unable to parse database configuration: %v", err)
	}

	poolConfig.MaxConns = 30
	poolConfig.MinConns = 3
	poolConfig.MaxConnLifetime = 30 * 60 // 30 minutes
	poolConfig.MaxConnIdleTime = 5 * 60  // 5 minutes
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
