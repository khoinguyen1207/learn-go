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
	var params dto.GetUsersRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	users, err := uh.service.GetUsers(ctx.Request.Context(), params.Search, params.OrderBy, params.Sort, params.Page, params.Limit)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	userDtos := dto.MapUsersToDto(users)

	dto.SuccessResponse(ctx, "Get users successfully", userDtos)
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	var params dto.GetUserByIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	dto.SuccessResponse(ctx, "Get user successfully", "")
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

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var uriReq dto.GetUserByIdRequest
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	var body dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	input := body.MapUpdateInputToModel(uriReq.ID)

	updatedUser, err := uh.service.UpdateUser(ctx.Request.Context(), input)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	data := dto.MapToUserDTO(updatedUser)

	dto.SuccessResponse(ctx, "User updated successfully", data)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var req dto.GetUserByUuidRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}

	_, err := uh.service.DeleteUser(ctx.Request.Context(), req.Uuid)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	dto.SuccessResponse(ctx, "User deleted successfully", true, http.StatusNoContent)
}

func (uh *UserHandler) RestoreUser(ctx *gin.Context) {
	var req dto.GetUserByUuidRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		dto.ValidationResponse(ctx, validation.HandleValidationError(err))
		return
	}
	_, err := uh.service.RestoreUser(ctx.Request.Context(), req.Uuid)
	if err != nil {
		dto.ErrorResponse(ctx, err)
		return
	}

	dto.SuccessResponse(ctx, "User restored successfully", true)
}
