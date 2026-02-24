package auth

import (
	"encoding/json"
	"project-shopping/internal/config"
	"project-shopping/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtService struct {
	cfg *config.Config
}

type EncryptedPayload struct {
	UserID string `json:"user_id"`
	Role   int32  `json:"role"`
}

type RefreshToken struct {
	Token     string    `json:"token"`
	UserId    string    `json:"user_id"`
	SessionId string    `json:"session_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewJWTService(cfg *config.Config) JWTService {
	return &jwtService{
		cfg: cfg,
	}
}

func (js *jwtService) GenerateAccessToken(user_id string, role int32) (string, error) {
	data := EncryptedPayload{
		UserID: user_id,
		Role:   role,
	}
	rawData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	encrypted, err := utils.EncryptAES(rawData, []byte(js.cfg.EncryptionKey))
	if err != nil {
		return "", err
	}

	expirationTime := parseExpirationTime(js.cfg.Jwt.AccessTokenExpiration, 15*time.Minute)
	claims := jwt.MapClaims{
		"data": encrypted,
		"type": "access",
		"jti":  uuid.New().String(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.cfg.Jwt.SecretKey))
}

func (js *jwtService) GenerateRefreshToken(user_id string, role int32) (RefreshToken, error) {
	expirationTime := parseExpirationTime(js.cfg.Jwt.RefreshTokenExpiration, 7*24*time.Hour)

	jti := uuid.New().String()
	expiresAt := time.Now().Add(expirationTime)

	claims := jwt.MapClaims{
		"user_id": user_id,
		"role":    role,
		"type":    "refresh",
		"jti":     jti,
		"iat":     time.Now().Unix(),
		"exp":     expiresAt.Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(js.cfg.Jwt.SecretKey))
	if err != nil {
		return RefreshToken{}, err
	}

	return RefreshToken{
		Token:     token,
		UserId:    user_id,
		SessionId: jti,
		ExpiresAt: expiresAt,
	}, nil
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

func (js *jwtService) VerifyAccessToken(token string) (*EncryptedPayload, error) {
	_, claims, err := js.VerifyToken(token)
	if err != nil {
		return nil, err
	}

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

func (js *jwtService) VerifyRefreshToken(token string) (*RefreshToken, error) {
	_, claims, err := js.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	userId, _ := claims["user_id"].(string)
	jti, _ := claims["jti"].(string)
	exp, _ := claims["exp"].(float64)
	if userId == "" || jti == "" || exp == 0 {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return &RefreshToken{
		UserId:    userId,
		SessionId: jti,
		ExpiresAt: time.Unix(int64(exp), 0),
	}, nil
}

func parseExpirationTime(durationStr string, defaultDuration time.Duration) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return defaultDuration
	}
	return duration
}
