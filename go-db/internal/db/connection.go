package db

import (
	"database/sql"
	"go-db/internal/configs"
)

var db *sql.DB

func InitDB(cfg configs.Config) error {
	dns := cfg.Db.ConnectionString
	var err error

	db, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}

	return nil
}
