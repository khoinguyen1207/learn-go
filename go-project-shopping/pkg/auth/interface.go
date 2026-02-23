package auth

import "github.com/golang-jwt/jwt/v5"

type JWTService interface {
	GenerateAccessToken(uuid, email string, role int32) (string, error)
	GenerateRefreshToken(uuid, email string, role int32) (string, error)
	VerifyToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error)
	DecryptAccessTokenPayload(claims jwt.MapClaims) (*EncryptedPayload, error)
}
