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

type CreateUserParams struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Password string `json:"password" binding:"required,strong_password"`
	Age      int    `json:"age" binding:"required,gt=0,lte=130"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UpdateUserParams struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Password string `json:"password" binding:"omitempty,strong_password"`
	Age      int    `json:"age" binding:"required,gt=0,lte=130"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

func (p *CreateUserParams) ToModel() model.User {
	return model.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
		Age:      p.Age,
		Status:   p.Status,
		Level:    p.Level,
	}
}

func (p *UpdateUserParams) ToModel() model.User {
	return model.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
		Age:      p.Age,
		Status:   p.Status,
		Level:    p.Level,
	}
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
		usersDto = append(usersDto, *MapToUserDTO(user))
	}

	return usersDto
}

func mapStatusText(status int) string {
	switch status {
	case 1:
		return "Active"
	case 2:
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
