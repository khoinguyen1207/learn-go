package main

import (
	"gin-project/internal/v1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			handler := handler.NewUserHandler()
			userGroup.GET("/:id", handler.GetUser)
			userGroup.GET("/slug/:slug", handler.GetUserBySlug)
			userGroup.POST("/", handler.CreateUser)
			userGroup.PUT("/:id", handler.UpdateUser)
			userGroup.DELETE("/:id", handler.DeleteUser)
		}
	}
	r.Run(":8080")
}
