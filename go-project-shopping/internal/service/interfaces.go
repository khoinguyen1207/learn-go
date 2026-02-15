package service

import (
	"context"
	"project-shopping/internal/db/sqlc"
)

type UserService interface {
	GetUsers(search string, page, limit int) error
	CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(id string) error
	UpdateUser(ctx context.Context, input sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, uuid string) (sqlc.User, error)
	RestoreUser(ctx context.Context, uuid string) (sqlc.User, error)
}
