package auth

import (
	"project-shopping/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfg *config.JwtConfig
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   int32  `json:"role"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

func NewJWTService(cfg *config.JwtConfig) JWTService {
	return &jwtService{
		cfg: cfg,
	}
}

func (js *jwtService) GenerateAccessToken(userID string, role int32) (string, error) {
	expirationTime := parseExpirationTime(js.cfg.AccessTokenExpiration, 15*time.Minute)
	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.cfg.SecretKey))
}

func (js *jwtService) GenerateRefreshToken(userID string, role int32) (string, error) {
	expirationTime := parseExpirationTime(js.cfg.RefreshTokenExpiration, 168*time.Hour)
	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.cfg.SecretKey))
}

func parseExpirationTime(durationStr string, defaultDuration time.Duration) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return defaultDuration
	}
	return duration
}
