package routes

import (
	"net/http"
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
	httpLogger := utils.NewLoggerWithPath("app.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("recovery.log", "error")
	rateLimitLogger := utils.NewLoggerWithPath("rate_limit.log", "warn")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Api-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

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

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
			"path":  ctx.Request.URL.Path,
		})
	})
}
