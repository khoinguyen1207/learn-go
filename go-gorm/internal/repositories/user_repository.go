package repositories

import (
	"database/sql"
	"go-gorm/internal/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id int) (models.User, error) {
	var user models.User

	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	row := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email)
	err := row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
