package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"

	"github.com/google/uuid"
)

type userRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FindAll() error {
	return nil
}

func (ur *userRepository) FindById(id string) bool {
	return false
}

func (ur *userRepository) Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	data, err := ur.db.CreateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	data, err := ur.db.UpdateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Delete(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	data, err := ur.db.SoftDeleteUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Restore(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	data, err := ur.db.RestoreUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) FindByEmail(email string) bool {
	return false
}
