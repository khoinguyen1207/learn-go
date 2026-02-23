package service

import (
	"context"
	"project-shopping/internal/db/sqlc"
)

type UserService interface {
	GetUsers(ctx context.Context, search string, orderBy, sort string, page, limit int32) ([]sqlc.User, int32, error)
	CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (sqlc.User, error)
	UpdateUser(ctx context.Context, input sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, uuid string) (sqlc.User, error)
	RestoreUser(ctx context.Context, uuid string) (sqlc.User, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password string) error
	Logout(ctx context.Context, token string) error
}
