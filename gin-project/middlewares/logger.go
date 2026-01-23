package middlewares

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(ctx *gin.Context) gin.HandlerFunc {
	logPath := "logs/app.log"

	err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return func(c *gin.Context) {
		logEvent := logger.Info()

		logEvent.Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("client_ip", c.ClientIP()).Str("query", c.Request.URL.RawQuery)

		c.Next()
	}
}
