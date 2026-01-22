package main

import (
	"gin-project/internal/v1/handler"
	"gin-project/middlewares"
	"gin-project/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

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
			categoryGroup.GET("/:category", middlewares.ApiKeyMiddleware(), handler.GetCategory)
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
			newsGroup.POST("/upload-multiple", handler.UploadMultipleNewsImages)
		}
	}

	r.StaticFS("/images", gin.Dir("./uploads", false))

	r.Run(":8080")
}
