package middleware

import (
	"fmt"
	"net/http"
	"project-shopping/internal/utils"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	recoveryLogger := utils.NewLoggerWithPath("internal/logs/recovery.log", "error")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				recoveryLogger.Error().
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
