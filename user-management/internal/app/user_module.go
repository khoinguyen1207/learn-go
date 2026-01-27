package app

import (
	"user-management/internal/handler"
	"user-management/internal/repository"
	"user-management/internal/routes"
	"user-management/internal/service"
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
