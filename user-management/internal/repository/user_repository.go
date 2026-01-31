package repository

import (
	"slices"
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

func (ur *userRepository) FindAll() ([]model.User, error) {
	return ur.users, nil
}

func (ur *userRepository) Create(user model.User) error {
	ur.users = append(ur.users, user)
	return nil
}

func (ur *userRepository) FindById(id string) (model.User, bool) {
	for _, user := range ur.users {
		if user.ID == id {
			return user, true
		}
	}
	return model.User{}, false
}

func (ur *userRepository) Update(id string, updatedUser model.User) error {
	for i, user := range ur.users {
		if user.ID == id {
			ur.users[i] = updatedUser
			return nil
		}
	}
	return nil
}

func (ur *userRepository) Delete(id string) error {
	index := slices.IndexFunc(ur.users, func(u model.User) bool {
		return u.ID == id
	})

	ur.users = slices.Delete(ur.users, index, index+1)
	return nil
}

func (ur *userRepository) FindByEmail(email string) (model.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return model.User{}, false
}
