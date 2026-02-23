package routes

import (
	"project-shopping/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	route := r.Group("/users")
	{
		route.GET("", ur.handler.GetUsers)
		route.POST("", ur.handler.CreateUser)
		route.GET("/:uuid", ur.handler.GetUserByUUID)
		route.PUT("/:id", ur.handler.UpdateUser)
		route.DELETE("/:uuid", ur.handler.DeleteUser)
		route.PUT("/restore/:uuid", ur.handler.RestoreUser)
	}
}
