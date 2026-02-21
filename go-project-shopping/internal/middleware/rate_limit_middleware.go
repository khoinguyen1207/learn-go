package middleware

import (
	"net/http"
	"project-shopping/internal/config"
	"project-shopping/internal/utils"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var clients = make(map[string]*Client)

func getRateLimiter(ip string) *rate.Limiter {
	client, exists := clients[ip]
	if !exists {
		RATE_LIMIT_REQUEST_SECOND := utils.GetEnvAsInt("RATE_LIMIT_REQUEST_SECOND", 5)
		RATE_LIMIT_REQUEST_BURST := utils.GetEnvAsInt("RATE_LIMIT_REQUEST_BURST", 10)
		// 5 requests per second with a burst of 10
		limiter := rate.NewLimiter(rate.Limit(RATE_LIMIT_REQUEST_SECOND), RATE_LIMIT_REQUEST_BURST)
		clients[ip] = &Client{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)

		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
	}
}

// ab -n 20 -c 1 -H "x-api-key: 21b8f79c-ba0e-485b-8a00-72b425a083a0" http://localhost:8080/api/v1/categories/books
func RateLimiterMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)
		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			if shouldLogRequest(ip) {
				logger.Warn().
					Str("method", c.Request.Method).
					Str("path", c.Request.URL.Path).
					Str("query", c.Request.URL.RawQuery).
					Str("client_ip", ip).
					Str("user_agent", c.Request.UserAgent()).
					Str("protocol", c.Request.Proto).
					Str("host", c.Request.Host).
					Str("remote_addr", c.Request.RemoteAddr).
					Interface("headers", c.Request.Header).
					Msg("Rate limit exceeded")
			}

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next()
	}
}

var rateLimitLogCache = sync.Map{}

func shouldLogRequest(ip string) bool {
	now := time.Now()

	last, exists := rateLimitLogCache.Load(ip)
	if exists {
		if now.Sub(last.(time.Time)) < config.RATE_LIMIT_LOG_TTL*time.Second {
			return false
		}
	}

	rateLimitLogCache.Store(ip, now)
	return true
}
