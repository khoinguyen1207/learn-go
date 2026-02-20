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
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", ur.handler.GetUsers)
		userGroup.POST("/", ur.handler.CreateUser)
		userGroup.GET("/:uuid", ur.handler.GetUserByUUID)
		userGroup.PUT("/:id", ur.handler.UpdateUser)
		userGroup.DELETE("/:uuid", ur.handler.DeleteUser)
		userGroup.PUT("/restore/:uuid", ur.handler.RestoreUser)
	}
}
