package routes

import (
	"fmt"
	"project-shopping/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	r.Use(
		middleware.RecoveryMiddleware(),
		middleware.LoggerMiddleware(),
		middleware.ApiKeyMiddleware(),
		middleware.RateLimiterMiddleware(),
	)

	apiGroup := r.Group("/api/v1")

	apiGroup.GET("/panic", func(c *gin.Context) {
		var a []int
		fmt.Println(a[1])
	})

	for _, route := range routes {
		route.Register(apiGroup)
	}
}
