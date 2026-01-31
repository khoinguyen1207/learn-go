package repository

import "user-management/internal/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	Create(user model.User) error
	FindById(id string) (model.User, bool)
	Update(id string, user model.User) error
	Delete(id string) error
	FindByEmail(email string) (model.User, bool)
}
