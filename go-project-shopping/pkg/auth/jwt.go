package auth

import (
	"encoding/json"
	"project-shopping/internal/config"
	"project-shopping/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfg *config.Config
}

type EncryptedPayload struct {
	UserID string `json:"user_id"`
	Role   int32  `json:"role"`
	Type   string `json:"type"`
}

func NewJWTService(cfg *config.Config) JWTService {
	return &jwtService{
		cfg: cfg,
	}
}

func (js *jwtService) GenerateAccessToken(uuid, email string, role int32) (string, error) {
	data := EncryptedPayload{
		UserID: uuid,
		Role:   role,
		Type:   "access",
	}
	rawData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	encrypted, err := utils.EncryptAES(rawData, []byte(js.cfg.EncryptionKey))
	if err != nil {
		return "", err
	}

	jwtConfig := js.cfg.Jwt
	expirationTime := parseExpirationTime(jwtConfig.AccessTokenExpiration, 15*time.Minute)
	claims := jwt.MapClaims{
		"data": encrypted,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SecretKey))
}

func (js *jwtService) GenerateRefreshToken(uuid, email string, role int32) (string, error) {
	return "", nil
}

func parseExpirationTime(durationStr string, defaultDuration time.Duration) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return defaultDuration
	}
	return duration
}
