package handler

import (
	"net/http"
	"user-management/internal/dto"
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
	var params dto.GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	users, err := uh.service.GetUsers(params.Search, params.Page, params.Limit)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	response.SuccessResponse(ctx, "Get users successfully", usersDto)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var params dto.CreateUserParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	user := params.ToModel()

	createdUser, err := uh.service.CreateUser(user)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	userDto := dto.MapToUserDTO(createdUser)

	response.SuccessResponse(ctx, "User created successfully", &userDto)
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	var params dto.GetUserByIdParams
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
	var uriParams dto.GetUserByIdParams
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	var body dto.UpdateUserParams
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	userToUpdate := body.ToModel()

	updatedUser, err := uh.service.UpdateUser(uriParams.ID, userToUpdate)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	userDto := dto.MapToUserDTO(updatedUser)

	response.SuccessResponse(ctx, "User updated successfully", &userDto)

}
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var params dto.GetUserByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	if err := uh.service.DeleteUser(params.ID); err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.SuccessResponse(ctx, "User deleted successfully", nil, http.StatusNoContent)
}
