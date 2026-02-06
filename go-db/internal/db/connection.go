package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-db/internal/configs"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *configs.Config) error {
	dns := cfg.Db.ConnectionString

	var err error
	DB, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}

	DB.SetMaxIdleConns(3)                   // Số kết nối nhàn rỗi tối đa
	DB.SetMaxOpenConns(30)                  // Số kết nối tối đa
	DB.SetConnMaxLifetime(30 * time.Minute) // Thời gian sống tối đa của một kết nối
	DB.SetConnMaxIdleTime(5 * time.Minute)  // Đóng kết nối nhàn rỗi sau khoảng thời gian này

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := DB.PingContext(context); err != nil {
		DB.Close()
		return fmt.Errorf("Could not connect to the database: %v", err)
	}

	return nil
}
