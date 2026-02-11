package handler

import (
	"net/http"
	"project-shopping/internal/dto"
	"project-shopping/internal/service"
	"project-shopping/internal/validation"

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
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	dto.SuccessResponse(ctx, "Get users successfully", "")
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	input := req.MapCreateInputToModel()

	createdUser, err := uh.service.CreateUser(ctx.Request.Context(), input)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	userDto := dto.MapToUserDTO(createdUser)

	dto.SuccessResponse(ctx, "User created successfully", userDto)
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	var params dto.GetUserByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	dto.SuccessResponse(ctx, "Get user successfully", "")
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var uriParams dto.GetUserByIdParams
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	var body dto.UpdateUserParams
	if err := ctx.ShouldBindJSON(&body); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	dto.SuccessResponse(ctx, "User updated successfully", "")
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var params dto.GetUserByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	dto.SuccessResponse(ctx, "User deleted successfully", "", http.StatusNoContent)
}
