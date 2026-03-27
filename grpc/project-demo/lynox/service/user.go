package service

import (
	"context"
	"lynox/model"
	"lynox/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
