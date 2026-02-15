package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() error
	FindById(id string) bool
	Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	Delete(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
	Restore(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
	FindByEmail(email string) bool
}
