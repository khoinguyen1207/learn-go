package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"user": name,
		})
	})

	r.GET("/search", func(ctx *gin.Context) {
		query := ctx.Query("q")
		ctx.JSON(200, gin.H{
			"query": query,
		})
	})

	r.Run(":8080")
}
