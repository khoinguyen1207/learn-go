package middlewares

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "logs/app.log"

	// Create log directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Open log file
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return func(c *gin.Context) {
		start := time.Now()

		logEvent := logger.Info()

		c.Next()

		status_code := c.Writer.Status()
		if status_code >= 500 {
			logEvent = logger.Error()
		} else if status_code >= 400 {
			logEvent = logger.Warn()
		}

		duration := time.Since(start)

		logEvent.Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("client_ip", c.ClientIP()).
			Str("query", c.Request.URL.RawQuery).
			Str("user_agent", c.Request.UserAgent()).
			Str("referer", c.Request.Referer()).
			Str("protocol", c.Request.Proto).
			Str("host", c.Request.Host).
			Str("remote_addr", c.Request.RemoteAddr).
			Str("request_uri", c.Request.RequestURI).
			Int64("content_length", c.Request.ContentLength).
			Interface("headers", c.Request.Header).
			Int("status", c.Writer.Status()).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request logged")
	}
}
