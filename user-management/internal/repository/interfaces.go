package repository

import "user-management/internal/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	Create(user model.User) error
	FindById(id string) (model.User, bool)
	Update()
	Delete()
	FindByEmail(email string) (model.User, bool)
}
