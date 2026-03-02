package main

import (
	"log"
	"path/filepath"
	"project-shopping/internal/config"
	"project-shopping/internal/utils"
	"project-shopping/pkg/logger"

	"github.com/joho/godotenv"
)

func NewWorker(config *config.Config) {}

func main() {
	// Load environment variables from .env file
	rootDir := utils.GetRootDir()
	filepath := filepath.Join(rootDir, ".env")

	if err := godotenv.Load(filepath); err != nil {
		log.Fatalf("⚠️ Error loading .env file: %v", err)
	} else {
		log.Println("✅ .env file loaded successfully")
	}

	// Initialize configuration
	config.NewConfig()

	// Init logger
	logger.InitLogger(config.Get().AppEnv)

	// Initialize worker
	NewWorker(config.Get())
}
