package service

import (
	"context"
	"project-shopping/internal/db/sqlc"
)

type UserService interface {
	GetUsers(search string, page, limit int) error
	CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(id string) error
	UpdateUser(id string) error
	DeleteUser(id string) error
}
