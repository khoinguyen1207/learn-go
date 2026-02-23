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

func (js *jwtService) VerifyToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(js.cfg.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, jwt.ErrTokenInvalidClaims
	}

	return token, claims, nil
}

func (js *jwtService) DecryptAccessTokenPayload(claims jwt.MapClaims) (*EncryptedPayload, error) {
	encryptedData, ok := (claims)["data"].(string)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	decrypted, err := utils.DecryptAES(encryptedData, []byte(js.cfg.EncryptionKey))
	if err != nil {
		return nil, err
	}

	var payload EncryptedPayload
	err = json.Unmarshal(decrypted, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func parseExpirationTime(durationStr string, defaultDuration time.Duration) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return defaultDuration
	}
	return duration
}
