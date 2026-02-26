package middleware

import (
	"context"
	"project-shopping/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.GetHeader("X-Trace-ID")
		if traceId == "" {
			traceId = uuid.New().String()
		}

		// Set trace ID in request context
		ctxValue := context.WithValue(ctx.Request.Context(), logger.TRACE_ID_KEY, traceId)
		ctx.Request = ctx.Request.WithContext(ctxValue)

		// Set trace ID in response header
		ctx.Writer.Header().Set("X-Trace-ID", traceId)

		// Set trace ID in gin context
		ctx.Set(logger.TRACE_ID_KEY, traceId)

		ctx.Next()
	}
}
