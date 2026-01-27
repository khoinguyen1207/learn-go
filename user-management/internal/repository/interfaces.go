package repository

import "user-management/internal/model"

type UserRepository interface {
	FindAll()
	Create(user model.User) error
	FindById()
	Update()
	Delete()
	FindByEmail(email string) (model.User, bool)
}
