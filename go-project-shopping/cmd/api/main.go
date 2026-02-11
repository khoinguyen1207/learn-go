package main

import (
	"log"
	"project-shopping/internal/app"
	"project-shopping/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// Run the application
	if err := application.Run(); err != nil {
		panic(err)
	}
}
