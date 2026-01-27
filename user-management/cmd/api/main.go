package main

import (
	"user-management/internal/app"
	"user-management/internal/config"
)

func main() {

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// Run the application
	if err := application.Run(); err != nil {
		panic(err)
	}
}
