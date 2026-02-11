package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"
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

func (ur *userRepository) Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	data, err := ur.db.CreateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) FindById(id string) bool {
	return false
}

func (ur *userRepository) Update(id string) error {
	return nil
}

func (ur *userRepository) Delete(id string) error {
	return nil
}

func (ur *userRepository) FindByEmail(email string) bool {
	return false
}
