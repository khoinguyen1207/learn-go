package routes

import (
	"user-management/internal/handler"

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
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", ur.handler.GetUsers)
		userGroup.POST("/", ur.handler.CreateUser)
		userGroup.GET("/:id", ur.handler.GetUserByID)
		userGroup.PUT("/:id", ur.handler.UpdateUser)
		userGroup.DELETE("/:id", ur.handler.DeleteUser)
	}
}
