package main

import (
	"user-management/internal/config"
	"user-management/internal/handler"
	"user-management/internal/repository"
	"user-management/internal/routes"
	"user-management/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	userRepository := repository.NewUserRepository()

	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	userRoutes := routes.NewUserRoutes(userHandler)

	r := gin.Default()

	routes.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.Port); err != nil {
		panic(err)
	}
}
