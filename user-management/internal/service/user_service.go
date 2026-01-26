package service

import (
	"log"
	"user-management/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetUsers() {
	log.Println("Service: GetUsers called")
	us.repo.FindAll()
}
func (us *userService) CreateUser() {

}
func (us *userService) GetUserByID() {

}
func (us *userService) UpdateUser() {

}
func (us *userService) DeleteUser() {

}
