package middleware

import (
	"net/http"
	"project-shopping/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService auth.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing token",
			})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		tokenPayload, err := jwtService.VerifyAccessToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		ctx.Set("user_id", tokenPayload.UserID)
		ctx.Set("role", tokenPayload.Role)

		ctx.Next()
	}
}
