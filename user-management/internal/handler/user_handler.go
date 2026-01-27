package handler

import (
	"net/http"
	"user-management/internal/model"
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
	uh.service.GetUsers()
}
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var body model.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, validation.HandleValidationError(err))
		return
	}

	createdUser, err := uh.service.CreateUser(body)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.SuccessResponse(ctx, "User created successfully", createdUser)
}
func (uh *UserHandler) GetUserByID(ctx *gin.Context) {

}
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {

}
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
