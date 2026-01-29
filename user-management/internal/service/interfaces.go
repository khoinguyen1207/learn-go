package service

import "user-management/internal/model"

type UserService interface {
	GetUsers(search string, page, limit int) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id string) (model.User, error)
	UpdateUser(id string, user model.User) (model.User, error)
	DeleteUser()
}
