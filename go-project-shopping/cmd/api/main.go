package main

import (
	"project-shopping/internal/app"
	"project-shopping/internal/config"
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
