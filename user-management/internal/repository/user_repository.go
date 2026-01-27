package repository

import (
	"log"
	"user-management/internal/model"
)

type userRepository struct {
	users []model.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []model.User{},
	}
}

func (ur *userRepository) FindAll() {
	log.Println("Repository: FindAll called")

}

func (ur *userRepository) Create(user model.User) error {
	ur.users = append(ur.users, user)
	return nil
}

func (ur *userRepository) FindById() {

}

func (ur *userRepository) Update() {

}

func (ur *userRepository) Delete() {

}

func (ur *userRepository) FindByEmail(email string) (model.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return model.User{}, false
}
