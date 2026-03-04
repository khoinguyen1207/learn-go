package app

import (
	"project-shopping/internal/handler"
	"project-shopping/internal/repository"
	"project-shopping/internal/routes"
	"project-shopping/internal/service"
	"project-shopping/pkg/mail"
	"project-shopping/pkg/rabbitmq"
)

type AuthModule struct {
	routes routes.Route
}

func NewAuthModule(mctx *ModuleContext, mailService mail.MailService, rabbitMQService rabbitmq.RabbitMQService) *AuthModule {
	userRepository := repository.NewUserRepository(mctx.db)
	authService := service.NewAuthService(userRepository, mctx.jwt, mctx.cache, mailService, rabbitMQService)
	authHandler := handler.NewAuthHandler(authService)
	authRoutes := routes.NewAuthRoutes(authHandler, mctx.jwt)

	return &AuthModule{
		routes: authRoutes,
	}
}

func (am *AuthModule) Routes() routes.Route {
	return am.routes
}
