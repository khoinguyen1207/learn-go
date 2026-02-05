package main

import (
	"go-db/internal/handlers"
	"go-db/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// cfg := configs.NewConfig()

	r := gin.Default()

	userRepository := repositories.NewUserRepository()
	userHandler := handlers.NewUserHandler(userRepository)

	r.Group("/api/v1")
	{
		r.GET("/user/:id", userHandler.GetUserById)
		r.POST("/user", userHandler.CreateUser)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
