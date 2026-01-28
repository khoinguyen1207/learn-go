package service

import "user-management/internal/model"

type UserService interface {
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id string) (model.User, error)
	UpdateUser()
	DeleteUser()
}
