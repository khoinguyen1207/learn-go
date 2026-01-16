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

type GetUserParams struct {
	ID string `uri:"id" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	var params GetUserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "GetUser called"})
}

func (h *UserHandler) GetUserBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if err := utils.RegexValidate("slug", slug, slugRegex); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
