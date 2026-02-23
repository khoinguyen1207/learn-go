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

		_, claims, err := jwtService.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		encryptedPayload, err := jwtService.DecryptAccessTokenPayload(claims)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		ctx.Set("user_id", encryptedPayload.UserID)
		ctx.Set("role", encryptedPayload.Role)

		ctx.Next()
	}
}
