package service

import (
	"context"
	"errors"
	"fmt"
	"project-shopping/internal/db/sqlc"
	"project-shopping/internal/repository"
	"project-shopping/internal/utils"
	"project-shopping/pkg/auth"
	"project-shopping/pkg/cache"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
)

type authService struct {
	repo   repository.UserRepository
	jwtSrv auth.JWTService
	cache  cache.CacheService
}

type LoginAttempt struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu               sync.Mutex
	clients          = make(map[string]*LoginAttempt)
	LoginAttemptTTL  = 5 * time.Minute
	MaxLoginAttempts = 5
)

func NewAuthService(repo repository.UserRepository, jwtSrv auth.JWTService, cache cache.CacheService) AuthService {
	return &authService{
		repo:   repo,
		jwtSrv: jwtSrv,
		cache:  cache,
	}
}

func (as *authService) getLoginAttempt(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {
		rateLimit := rate.Limit(float64(MaxLoginAttempts) / LoginAttemptTTL.Seconds())
		limiter := rate.NewLimiter(rateLimit, MaxLoginAttempts)
		clients[ip] = &LoginAttempt{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func (as *authService) getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func (as *authService) checkLoginAttempts(ip string) error {
	limiter := as.getLoginAttempt(ip)
	if !limiter.Allow() {
		return errors.New("Too many login attempts. Please try again later.")
	}
	return nil
}

func (as *authService) cleanupClients(ip string) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, ip)
}

func (as *authService) Login(context *gin.Context, email, password string) (string, string, error) {
	ctx := context.Request.Context()
	ip := as.getClientIP(context)
	if err := as.checkLoginAttempts(ip); err != nil {
		return "", "", utils.NewError(err.Error(), utils.CodeTooManyRequests)
	}

	email = utils.NormalizeString(email)
	user, err := as.repo.FindByEmail(ctx, email)
	if err != nil {
		as.getLoginAttempt(ip)
		return "", "", utils.NewError("Invalid email or password", utils.CodeUnauthorized)
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if bcryptErr != nil {
		as.getLoginAttempt(ip)
		return "", "", utils.NewError("Invalid email or password", utils.CodeUnauthorized)
	}

	accessToken, err := as.jwtSrv.GenerateAccessToken(user.Uuid.String(), user.Level)
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to generate access token", utils.CodeInternalServerError)
	}

	refreshToken, err := as.jwtSrv.GenerateRefreshToken(user.Uuid.String(), user.Level)
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to generate refresh token", utils.CodeInternalServerError)
	}

	key := fmt.Sprintf("session:%s", refreshToken.SessionId)
	err = as.cache.Set(ctx, key, refreshToken, time.Until(refreshToken.ExpiresAt))
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to store refresh token", utils.CodeInternalServerError)
	}

	// Clean up old login attempts for this IP
	as.cleanupClients(ip)

	return accessToken, refreshToken.Token, nil
}

func (as *authService) RefreshToken(ctx context.Context, token string) (string, string, error) {
	// Verify the refresh token
	tokenPayload, err := as.jwtSrv.VerifyRefreshToken(token)
	if err != nil {
		return "", "", utils.NewError("Token invalid", utils.CodeUnauthorized)
	}

	// Check if the refresh token exists
	cacheKey := fmt.Sprintf("session:%s", tokenPayload.SessionId)
	exists, err := as.cache.Exists(ctx, cacheKey)
	if err != nil || !exists {
		return "", "", utils.NewError("Token invalid or expired", utils.CodeUnauthorized)
	}

	userUuid, _ := uuid.Parse(tokenPayload.UserId)
	user, err := as.repo.FindByUUID(ctx, userUuid)
	if err != nil {
		return "", "", utils.NewError("User not found", utils.CodeUnauthorized)
	}

	// Revoke the old token
	as.cache.Delete(ctx, cacheKey)

	// Generate new access token
	accessToken, err := as.jwtSrv.GenerateAccessToken(tokenPayload.UserId, user.Level)
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to generate access token", utils.CodeInternalServerError)
	}

	// Generate new refresh token
	newRefreshToken, err := as.jwtSrv.GenerateRefreshToken(tokenPayload.UserId, user.Level)
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to generate refresh token", utils.CodeInternalServerError)
	}

	cacheKey = fmt.Sprintf("session:%s", newRefreshToken.SessionId)
	err = as.cache.Set(ctx, cacheKey, newRefreshToken, time.Until(newRefreshToken.ExpiresAt))
	if err != nil {
		return "", "", utils.WrapError(err, "Failed to store refresh token", utils.CodeInternalServerError)
	}

	return accessToken, newRefreshToken.Token, nil
}

func (as *authService) Register(ctx context.Context, email, password string) error {
	return nil
}

func (as *authService) Logout(ctx context.Context, refreshToken string) error {
	tokenPayload, err := as.jwtSrv.VerifyRefreshToken(refreshToken)
	if err != nil {
		return utils.NewError("Token invalid", utils.CodeUnauthorized)
	}

	cacheKey := fmt.Sprintf("session:%s", tokenPayload.SessionId)
	err = as.cache.Delete(ctx, cacheKey)
	if err != nil {
		return utils.WrapError(err, "Failed to revoke token", utils.CodeInternalServerError)
	}

	return nil
}

func (as *authService) ForgotPassword(ctx context.Context, email string) error {
	rateLimitKey := fmt.Sprintf("reset:rate:%s", email)

	if exists, err := as.cache.Exists(ctx, rateLimitKey); err != nil {
		return utils.WrapError(err, "Failed to check rate limit", utils.CodeInternalServerError)
	} else if exists {
		return utils.NewError("Too many requests. Please try again later.", utils.CodeTooManyRequests)
	}

	user, err := as.repo.FindByEmail(ctx, email)
	if err != nil {
		return utils.NewError("User not found", utils.CodeNotFound)
	}

	token, err := utils.GenerateRandomString(16)
	if err != nil {
		return utils.WrapError(err, "Failed to generate reset token", utils.CodeInternalServerError)
	}

	cacheKey := fmt.Sprintf("reset:%s", token)
	if err = as.cache.Set(ctx, cacheKey, user.Uuid.String(), 15*time.Minute); err != nil {
		return utils.WrapError(err, "Failed to store reset token", utils.CodeInternalServerError)
	}

	if err = as.cache.Set(ctx, rateLimitKey, "1", 5*time.Minute); err != nil {
		return utils.WrapError(err, "Failed to set rate limit", utils.CodeInternalServerError)
	}

	resetLink := fmt.Sprintf("https://yourdomain.com/reset-password?token=%s", token)
	// Simulate sending email
	fmt.Printf("Sending password reset link to %s: %s\n", email, resetLink)

	return nil
}

func (as *authService) ResetPassword(ctx context.Context, token, newPassword string) error {

	cacheKey := fmt.Sprintf("reset:%s", token)
	var uuidStr string
	err := as.cache.Get(ctx, cacheKey, &uuidStr)
	if err == redis.Nil || uuidStr == "" {
		return utils.NewError("Invalid or expired token", utils.CodeBadRequest)
	}
	if err != nil {
		return utils.WrapError(err, "Failed to get reset token", utils.CodeInternalServerError)
	}

	userUuid, err := uuid.Parse(uuidStr)
	if err != nil {
		return utils.WrapError(err, "Failed to parse user UUID", utils.CodeInternalServerError)
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return utils.WrapError(err, "Failed to hash password", utils.CodeInternalServerError)
	}

	arg := sqlc.UpdateUserPasswordParams{
		Uuid:     userUuid,
		Password: string(hashedPassword),
	}
	err = as.repo.UpdatePassword(ctx, arg)
	if err != nil {
		return utils.WrapError(err, "Failed to update password", utils.CodeInternalServerError)
	}

	as.cache.Delete(ctx, cacheKey)

	return nil
}
