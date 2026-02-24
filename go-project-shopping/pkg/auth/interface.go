package auth

import "github.com/golang-jwt/jwt/v5"

type JWTService interface {
	GenerateAccessToken(user_id string, role int32) (string, error)
	GenerateRefreshToken(user_id string, role int32) (RefreshToken, error)
	VerifyToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error)
	DecryptAccessTokenPayload(claims jwt.MapClaims) (*EncryptedPayload, error)
}
