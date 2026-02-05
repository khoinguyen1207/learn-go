package handlers

import (
	"go-db/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ur repositories.UserRepository
}

func NewUserHandler(ur repositories.UserRepository) *UserHandler {
	return &UserHandler{
		ur: ur,
	}
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	// Implementation for getting a user
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {

}
