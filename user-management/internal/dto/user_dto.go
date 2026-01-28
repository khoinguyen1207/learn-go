package dto

import "user-management/internal/model"

type UserDTO struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"full_name"`
	Age    int    `json:"age"`
	Status string `json:"status"`
	Level  string `json:"level"`
}

func MapToUserDTO(user model.User) *UserDTO {
	return &UserDTO{
		Id:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Age:    user.Age,
		Status: mapStatusText(user.Status),
		Level:  mapLevelText(user.Level),
	}
}

func MapUsersToDto(users []model.User) []UserDTO {
	usersDto := make([]UserDTO, 0, len(users))

	for _, user := range users {
		userDto := UserDTO{
			Id:     user.ID,
			Email:  user.Email,
			Name:   user.Name,
			Age:    user.Age,
			Status: mapStatusText(user.Status),
			Level:  mapLevelText(user.Level),
		}
		usersDto = append(usersDto, userDto)
	}

	return usersDto
}

func mapStatusText(status int) string {
	switch status {
	case 1:
		return "Active"
	case 0:
		return "Inactive"
	default:
		return "Unknown"
	}
}

func mapLevelText(level int) string {
	switch level {
	case 1:
		return "Admin"
	case 2:
		return "User"
	default:
		return "Guest"
	}
}
