package repository

import (
	"context"
	"lynox/model"
)

type User struct {
	ID   int64
	Name string
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	// mock data
	return &model.User{
		ID:      id,
		Name:    "Nguyen",
		Balance: 1000,
	}, nil
}
