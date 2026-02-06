package main

import (
	"go-gorm/internal/configs"
	"go-gorm/internal/db"
	"go-gorm/internal/handlers"
	"go-gorm/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	cfg := configs.NewConfig()
	if err := db.InitDB(cfg); err != nil {
		panic(err)
	}

	r := gin.Default()

	userRepository := repositories.NewUserRepository(db.DB)
	userHandler := handlers.NewUserHandler(userRepository)

	api := r.Group("/api/v1")
	{
		api.GET("/users/:id", userHandler.GetUserById)
		api.POST("/users", userHandler.CreateUser)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
