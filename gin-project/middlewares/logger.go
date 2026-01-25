package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
		request_body := make(map[string]any)
		formFiles := []map[string]any{}

		contentType := c.GetHeader("Content-Type")
		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
				logger.Error().Err(err).Msg("Failed to parse multipart form")
			} else {
				// Extract form values
				for key, vals := range c.Request.MultipartForm.Value {
					if len(vals) == 1 {
						request_body[key] = vals[0]
						continue
					}
					request_body[key] = vals
				}

				// Extract file names
				for key, files := range c.Request.MultipartForm.File {
					for _, f := range files {
						formFiles = append(formFiles, map[string]any{
							"form_field": key,
							"file_name":  f.Filename,
							"file_size":  f.Size,
							"mime_type":  f.Header.Get("Content-Type"),
						})
					}
				}
				if len(formFiles) > 0 {
					request_body["files"] = formFiles
				}
			}
		} else {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to read request body")
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			if strings.HasPrefix(contentType, "application/json") {
				_ = json.Unmarshal(bodyBytes, &request_body)
			} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
				values, _ := url.ParseQuery(string(bodyBytes))
				for key, vals := range values {
					if len(vals) == 1 {
						request_body[key] = vals[0]
						continue
					}
					request_body[key] = vals

				}
			}
		}

		c.Next()

		status_code := c.Writer.Status()
		logEvent := logger.Info()
		if status_code >= 500 {
			logEvent = logger.Error()
		} else if status_code >= 400 {
			logEvent = logger.Warn()
		}

		duration := time.Since(start)

		logEvent.Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("host", c.Request.Host).
			Str("request_uri", c.Request.RequestURI).
			Str("content-type", c.GetHeader("Content-Type")).
			Interface("headers", c.Request.Header).
			Interface("body", request_body).
			Int("status", c.Writer.Status()).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request logged")
	}
}
