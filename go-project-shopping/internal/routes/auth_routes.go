package routes

import (
	"project-shopping/internal/handler"
	"project-shopping/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	handler    *handler.AuthHandler
	jwtService auth.JWTService
}

func NewAuthRoutes(handler *handler.AuthHandler, jwtService auth.JWTService) *AuthRoutes {
	return &AuthRoutes{
		handler:    handler,
		jwtService: jwtService,
	}
}

func (ur *AuthRoutes) Register(r *gin.RouterGroup) {
	route := r.Group("/auth")
	{
		route.POST("/login", ur.handler.Login)
		route.POST("/register", ur.handler.Register)
		route.POST("/logout", ur.handler.Logout)
		route.POST("/refresh-token", ur.handler.RefreshToken)
	}
}
