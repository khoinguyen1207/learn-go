package repositories

import "go-gorm/internal/models"

type UserRepository interface {
	FindById(id int) (models.User, error)
	CreateUser(user *models.User) error
}
