package routes

import (
	"project-shopping/internal/middleware"
	"project-shopping/internal/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	httpLogger := utils.NewLoggerWithPath("internal/logs/app.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("internal/logs/recovery.log", "error")
	rateLimitLogger := utils.NewLoggerWithPath("internal/logs/rate_limit.log", "warn")

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
