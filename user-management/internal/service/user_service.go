package service

import (
	"log"
	"user-management/internal/model"
	"user-management/internal/repository"
	"user-management/internal/utils"

	"github.com/google/uuid"
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
func (us *userService) CreateUser(user model.User) (model.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return model.User{}, utils.NewError("Email already exists", utils.CodeBadRequest)
	}

	user.ID = uuid.New().String()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return model.User{}, utils.WrapError(err, "Failed to hash password", utils.CodeInternalServerError)
	}
	user.Password = string(hashedPassword)

	if err := us.repo.Create(user); err != nil {
		return model.User{}, utils.WrapError(err, "Failed to create user", utils.CodeInternalServerError)
	}

	return user, nil
}
func (us *userService) GetUserByID() {

}
func (us *userService) UpdateUser() {

}
func (us *userService) DeleteUser() {

}
