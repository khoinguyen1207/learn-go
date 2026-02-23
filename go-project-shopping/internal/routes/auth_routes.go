package routes

import (
	"project-shopping/internal/handler"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	handler *handler.AuthHandler
}

func NewAuthRoutes(handler *handler.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
	}
}

func (ur *AuthRoutes) Register(r *gin.RouterGroup) {
	route := r.Group("/auth")
	{
		route.POST("/login", ur.handler.Login)
		route.POST("/register", ur.handler.Register)
		route.POST("/logout", ur.handler.Logout)
	}
}
