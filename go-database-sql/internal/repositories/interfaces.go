package repositories

import "go-database-sql/internal/models"

type UserRepository interface {
	FindById(id int) (models.User, error)
	CreateUser(user *models.User) error
}
