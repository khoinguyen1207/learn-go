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

func NewUserModule(mctx *ModuleContext) *UserModule {
	userRepository := repository.NewUserRepository(mctx.db)
	userService := service.NewUserService(userRepository, mctx.cache)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler, mctx.jwt)

	return &UserModule{
		routes: userRoutes,
	}
}

func (um *UserModule) Routes() routes.Route {
	return um.routes
}
