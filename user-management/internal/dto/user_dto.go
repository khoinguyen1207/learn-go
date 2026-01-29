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

type GetUserByIdParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=100"`
	Page   int    `form:"page" binding:"omitempty,gte=1"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

type UpdateUserParams struct {
	Name     string `json:"name" binding:"omitempty,min=3,max=100"`
	Email    string `json:"email" binding:"omitempty,email,email_advanced"`
	Password string `json:"password" binding:"omitempty,strong_password"`
	Age      int    `json:"age" binding:"omitempty,gt=0,lte=130"`
	Status   int    `json:"status" binding:"omitempty,oneof=1 2"`
	Level    int    `json:"level" binding:"omitempty,oneof=1 2"`
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
