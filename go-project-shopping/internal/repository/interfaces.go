package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"
)

type UserRepository interface {
	FindAll() error
	FindById(id string) bool
	Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	Delete(id string) error
	FindByEmail(email string) bool
}
