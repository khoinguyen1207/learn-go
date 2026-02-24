package service

import (
	"context"
	"fmt"
	"project-shopping/internal/repository"
	"project-shopping/internal/utils"
	"project-shopping/pkg/auth"
	"project-shopping/pkg/cache"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo   repository.UserRepository
	jwtSrv auth.JWTService
	cache  cache.CacheService
}

func NewAuthService(repo repository.UserRepository, jwtSrv auth.JWTService, cache cache.CacheService) AuthService {
	return &authService{
		repo:   repo,
		jwtSrv: jwtSrv,
		cache:  cache,
	}
}

func (as *authService) Login(ctx context.Context, email, password string) (string, string, error) {
	email = utils.NormalizeString(email)
	user, err := as.repo.FindByEmail(ctx, email)

	if err != nil {
		return "", "", utils.NewError("Invalid email or password", utils.CodeUnauthorized)
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if bcryptErr != nil {
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
