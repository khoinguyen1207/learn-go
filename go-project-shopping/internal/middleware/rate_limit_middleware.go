package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
		limiter := rate.NewLimiter(5, 10) // 5 requests per second with a burst of 10
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
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		ctx.Next()
	}
}
