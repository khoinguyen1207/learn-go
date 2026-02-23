package service

import (
	"context"
	"project-shopping/internal/repository"
	"project-shopping/internal/utils"
	"project-shopping/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo   repository.UserRepository
	jwtSrv auth.JWTService
}

func NewAuthService(repo repository.UserRepository, jwtSrv auth.JWTService) AuthService {
	return &authService{
		repo:   repo,
		jwtSrv: jwtSrv,
	}
}

func (as *authService) Login(ctx context.Context, email, password string) (string, error) {
	email = utils.NormalizeString(email)
	user, err := as.repo.FindByEmail(ctx, email)

	if err != nil {
		return "", utils.NewError("Invalid email or password", utils.CodeUnauthorized)
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if bcryptErr != nil {
		return "", utils.NewError("Invalid email or password", utils.CodeUnauthorized)
	}

	token, err := as.jwtSrv.GenerateAccessToken(user.Uuid.String(), user.Email, user.Level)
	if err != nil {
		return "", utils.WrapError(err, "Failed to generate access token", utils.CodeInternalServerError)
	}

	return token, nil
}

func (as *authService) Register(ctx context.Context, email, password string) error {
	return nil
}

func (as *authService) Logout(ctx context.Context, token string) error {
	return nil
}
