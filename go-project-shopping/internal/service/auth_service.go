package service

import (
	"context"
	"fmt"
	"project-shopping/internal/repository"
	"project-shopping/internal/utils"
	"project-shopping/pkg/auth"
	"project-shopping/pkg/cache"
	"time"

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

func (as *authService) Register(ctx context.Context, email, password string) error {
	return nil
}

func (as *authService) Logout(ctx context.Context, token string) error {
	return nil
}
