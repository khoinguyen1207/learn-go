package service

import (
	"project-shopping/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetUsers(search string, page, limit int) error {
	return nil
}

func (us *userService) CreateUser() error {
	return nil
}

func (us *userService) GetUserByID(id string) error {
	return nil
}

func (us *userService) UpdateUser(id string) error {
	return nil
}

func (us *userService) DeleteUser(id string) error {
	return nil
}
