package routes

import (
	"project-shopping/internal/middleware"
	"project-shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	httpLogger := utils.NewLoggerWithPath("internal/logs/app.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("internal/logs/recovery.log", "error")
	rateLimitLogger := utils.NewLoggerWithPath("internal/logs/rate_limit.log", "warn")

	r.Use(
		middleware.RateLimiterMiddleware(rateLimitLogger),
		middleware.TraceMiddleware(),
		middleware.RecoveryMiddleware(recoveryLogger),
		middleware.LoggerMiddleware(httpLogger),
		middleware.ApiKeyMiddleware(),
	)

	apiGroup := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(apiGroup)
	}
}
