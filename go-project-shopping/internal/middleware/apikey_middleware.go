package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	apiKey := os.Getenv("X_API_KEY")

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

		if clientApiKey != apiKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			return
		}

		ctx.Next()
	}
}
