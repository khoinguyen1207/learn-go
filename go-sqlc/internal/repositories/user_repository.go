package repositories

import (
	"context"
	"go-sqlc/internal/db/sqlc"

	"github.com/google/uuid"
)

type userRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByUuid(ctx context.Context, id uuid.UUID) (sqlc.User, error) {
	user, err := ur.db.GetUserByUuid(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := ur.db.CreateUser(ctx, input)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
