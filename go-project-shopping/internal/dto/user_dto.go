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
	Age       *int32 `json:"age"`
	Status    string `json:"status"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserByIdRequest struct {
	ID int `uri:"id" binding:"required,gte=0"`
}

type GetUserByUuidRequest struct {
	Uuid string `uri:"uuid" binding:"required,uuid"`
}

type GetUsersRequest struct {
	Search string `form:"search" binding:"omitempty,min=3,max=100"`
	Page   int    `form:"page" binding:"omitempty,gte=1"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Password string `json:"password" binding:"required,strong_password"`
	Age      *int32 `json:"age" binding:"omitempty,gt=0,lte=120"`
	Status   int32  `json:"status" binding:"required,oneof=1 2 3"`
	Level    int32  `json:"level" binding:"required,oneof=1 2 3"`
}

type UpdateUserRequest struct {
	Fullname *string `json:"fullname" binding:"omitempty,min=3,max=100"`
	Password *string `json:"password" binding:"omitempty,strong_password"`
	Age      *int32  `json:"age" binding:"omitempty,gt=0,lte=120"`
	Status   *int32  `json:"status" binding:"omitempty,oneof=1 2 3"`
	Level    *int32  `json:"level" binding:"omitempty,oneof=1 2 3"`
}

func (input *CreateUserRequest) MapCreateInputToModel() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
		Age:      input.Age,
		Status:   input.Status,
		Level:    input.Level,
	}
}

func (input *UpdateUserRequest) MapUpdateInputToModel(id int) sqlc.UpdateUserParams {
	id32 := int32(id)
	return sqlc.UpdateUserParams{
		ID:       &id32,
		Fullname: input.Fullname,
		Password: input.Password,
		Age:      input.Age,
		Status:   input.Status,
		Level:    input.Level,
	}
}

func MapToUserDTO(user sqlc.User) *UserDTO {
	userDto := &UserDTO{
		Id:        int(user.ID),
		Uuid:      user.Uuid.String(),
		Fullname:  user.Fullname,
		Email:     user.Email,
		Age:       user.Age,
		Status:    mapStatusText(int(user.Status)),
		Level:     mapLevelText(int(user.Level)),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
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
