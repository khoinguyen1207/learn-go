package service

import (
	"strings"
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

func (us *userService) GetUsers(search string, page, limit int) ([]model.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(err, "Failed to get users", utils.CodeInternalServerError)
	}

	filteredUsers := []model.User{}

	if search != "" {
		for _, user := range users {
			email := utils.NormalizeString(user.Email)
			name := utils.NormalizeString(user.Name)

			if strings.Contains(email, search) || strings.Contains(name, search) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	} else {
		filteredUsers = users
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(filteredUsers) {
		return []model.User{}, nil
	}
	if end > len(filteredUsers) {
		end = len(filteredUsers)
	}

	return filteredUsers[start:end], nil
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

func (us *userService) GetUserByID(id string) (model.User, error) {
	user, exist := us.repo.FindById(id)
	if !exist {
		return model.User{}, utils.NewError("User not found", utils.CodeNotFound)
	}
	return user, nil
}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
