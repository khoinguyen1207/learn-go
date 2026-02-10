package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RecoveryMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				logger.Error().
					Str("path", c.Request.URL.Path).
					Str("method", c.Request.Method).
					Str("client_ip", c.ClientIP()).
					Str("error", fmt.Sprintf("%v", err)).
					Str("stack", string(stack)).
					Msg("Panic recovered")

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
					"error":   fmt.Sprintf("%v", err),
				})
			}
		}()

		c.Next()
	}
}
