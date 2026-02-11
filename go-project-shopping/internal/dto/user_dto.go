package dto

import (
	"project-shopping/internal/db/sqlc"
	"time"
)

type UserDTO struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Email     string `json:"email"`
	Fullname  string `json:"full_name"`
	Age       *int   `json:"age"`
	Status    string `json:"status"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserByIdParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=100"`
	Page   int    `form:"page" binding:"omitempty,gte=1"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Password string `json:"password" binding:"required,strong_password"`
	Age      *int   `json:"age" binding:"omitempty,gt=0,lte=120"`
	Status   int    `json:"status" binding:"required,oneof=1 2 3"`
	Level    int    `json:"level" binding:"required,oneof=1 2 3"`
}

type UpdateUserParams struct {
	Fullname string `json:"fullname" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Password string `json:"password" binding:"omitempty,strong_password"`
	Age      int    `json:"age" binding:"required,gt=0,lte=120"`
	Status   int    `json:"status" binding:"required,oneof=1 2 3"`
	Level    int    `json:"level" binding:"required,oneof=1 2 3"`
}

func (input *CreateUserRequest) MapCreateInputToModel() sqlc.CreateUserParams {
	var int32Age *int32

	if input.Age != nil {
		age := int32(*input.Age)
		int32Age = &age
	}

	return sqlc.CreateUserParams{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
		Age:      int32Age,
		Status:   int32(input.Status),
		Level:    int32(input.Level),
	}
}

func MapToUserDTO(user sqlc.User) *UserDTO {
	userDto := &UserDTO{
		Id:        int(user.ID),
		Uuid:      user.Uuid.String(),
		Fullname:  user.Fullname,
		Email:     user.Email,
		Status:    mapStatusText(int(user.Status)),
		Level:     mapLevelText(int(user.Level)),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	if user.Age != nil {
		age := int(*user.Age)
		userDto.Age = &age
	}

	return userDto
}

func MapUsersToDto() []UserDTO {
	return []UserDTO{}
}

func mapStatusText(status int) string {
	switch status {
	case 1:
		return "Active"
	case 2:
		return "Inactive"
	case 3:
		return "Banned"
	default:
		return "Unknown"
	}
}

func mapLevelText(level int) string {
	switch level {
	case 1:
		return "Administrator"
	case 2:
		return "Moderator"
	case 3:
		return "Member"
	default:
		return "Unknown"
	}
}
