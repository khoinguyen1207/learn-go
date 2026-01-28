package handler

import (
	"user-management/internal/dto"
	"user-management/internal/model"
	"user-management/internal/request"
	"user-management/internal/response"
	"user-management/internal/service"
	"user-management/internal/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.service.GetUsers()
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	response.SuccessResponse(ctx, "Get users successfully", usersDto)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var body model.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	createdUser, err := uh.service.CreateUser(body)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	userDto := dto.MapToUserDTO(createdUser)

	response.SuccessResponse(ctx, "User created successfully", &userDto)
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	var params request.GetUserByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	user, err := uh.service.GetUserByID(params.ID)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	userDto := dto.MapToUserDTO(user)

	response.SuccessResponse(ctx, "Get user successfully", &userDto)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
