package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	apiGroup := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(apiGroup)
	}
}
