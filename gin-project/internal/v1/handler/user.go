package handler

import (
	"gin-project/internal/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var slugRegex = regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if !utils.IsValidID(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GetUser called"})
}

func (h *UserHandler) GetUserBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slug format"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GetUserBySlug called"})
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "CreateUser called"})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UpdateUser called"})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "DeleteUser called"})
}
