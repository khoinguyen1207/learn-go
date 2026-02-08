package app

import (
	"project-shopping/internal/handler"
	"project-shopping/internal/repository"
	"project-shopping/internal/routes"
	"project-shopping/internal/service"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (um *UserModule) Routes() routes.Route {
	return um.routes
}
