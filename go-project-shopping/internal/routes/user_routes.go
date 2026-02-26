package routes

import (
	"project-shopping/internal/handler"
	"project-shopping/internal/middleware"
	"project-shopping/pkg/auth"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler    *handler.UserHandler
	jwtService auth.JWTService
}

func NewUserRoutes(handler *handler.UserHandler, jwtService auth.JWTService) *UserRoutes {
	return &UserRoutes{
		handler:    handler,
		jwtService: jwtService,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	route := r.Group("/users")

	{
		route.GET("", ur.handler.GetUsers)
		route.POST("", ur.handler.CreateUser)
		route.Use(middleware.AuthMiddleware(ur.jwtService))
		route.GET("/:uuid", ur.handler.GetUserByUUID)
		route.PUT("/:id", ur.handler.UpdateUser)
		route.DELETE("/:uuid", ur.handler.DeleteUser)
		route.PUT("/restore/:uuid", ur.handler.RestoreUser)
	}
}
