package main

import (
	"gin-project/internal/utils"
	"gin-project/internal/v1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := utils.RegisterCustomValidations(); err != nil {
		panic(err)
	}

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

		categoryGroup := v1.Group("/categories")
		{
			handler := handler.NewCategoryHandler()
			categoryGroup.GET("/:category", handler.GetCategory)
		}

		productGroup := v1.Group("/products")
		{
			handler := handler.NewProductHandler()
			productGroup.GET("/", handler.GetProduct)
			productGroup.POST("/", handler.CreateProduct)
		}

		newsGroup := v1.Group("/news")
		{
			handler := handler.NewNewsHandler()
			newsGroup.POST("/", handler.CreateNews)
			newsGroup.POST("/upload", handler.UploadNewsImage)
		}
	}
	r.Run(":8080")
}
