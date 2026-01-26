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

func (ur *userRepository) Create() {

}

func (ur *userRepository) FindById() {

}

func (ur *userRepository) Update() {

}

func (ur *userRepository) Delete() {

}
