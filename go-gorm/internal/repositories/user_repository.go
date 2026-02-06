package repositories

import (
	"go-gorm/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id int) (models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
