package service

import "user-management/internal/model"

type UserService interface {
	GetUsers()
	CreateUser(user model.User) (model.User, error)
	GetUserByID()
	UpdateUser()
	DeleteUser()
}
