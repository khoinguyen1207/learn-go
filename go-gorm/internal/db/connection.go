package db

import (
	"context"
	"fmt"
	"go-gorm/internal/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(cfg *configs.Config) error {
	dns := cfg.Db.ConnectionString

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("Could not get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(3)                   // Số kết nối nhàn rỗi tối đa
	sqlDB.SetMaxOpenConns(30)                  // Số kết nối tối đa
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Thời gian sống tối đa của một kết nối
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // Đóng kết nối nhàn rỗi sau khoảng thời gian này

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return fmt.Errorf("Could not connect to the database: %v", err)
	}

	return nil
}
