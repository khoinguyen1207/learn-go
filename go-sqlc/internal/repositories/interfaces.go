package repositories

import (
	"context"
	"go-sqlc/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByUuid(ctx context.Context, id uuid.UUID) (sqlc.User, error)
	CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
}
