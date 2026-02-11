package main

import (
	"log"
	"project-shopping/internal/app"
	"project-shopping/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// Run the application
	if err := application.Run(); err != nil {
		log.Fatalf("Application run error: %v", err)
	}
}
