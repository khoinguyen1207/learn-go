package main

import (
	"log"
	"os"
	"path/filepath"
	"project-shopping/internal/app"
	"project-shopping/internal/config"
	"project-shopping/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	loadEnv()

	// Initialize configuration
	config.NewConfig()

	// Init logger
	logger.InitLogger(config.Get().AppEnv)

	// Initialize application
	application := app.NewApplication(config.Get())

	// Run the application
	if err := application.Run(); err != nil {
		log.Fatalf("Application run error: %v", err)
	}
}

func loadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Error getting current working directory: %v", err)
	}
	filepath := filepath.Join(dir, ".env")

	if err := godotenv.Load(filepath); err != nil {
		log.Fatalf("⚠️ Error loading .env file: %v", err)
	} else {
		log.Println("✅ .env file loaded successfully")
	}
}
