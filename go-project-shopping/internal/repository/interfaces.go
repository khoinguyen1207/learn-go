package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"
)

type UserRepository interface {
	FindAll() error
	Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	FindById(id string) bool
	Update(id string) error
	Delete(id string) error
	FindByEmail(email string) bool
}
