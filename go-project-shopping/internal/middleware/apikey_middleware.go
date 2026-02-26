package middleware

import (
	"net/http"
	"project-shopping/internal/config"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	cfg := config.Get()

	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}

		clientApiKey := ctx.GetHeader("x-api-key")
		if clientApiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing API key"})
			return
		}

		if clientApiKey != cfg.XApiKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			return
		}

		ctx.Next()
	}
}
