package handlers

import (
	"database/sql"
	"errors"
	"go-db/internal/models"
	"go-db/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uh.repo.FindById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Fetched user successfully", "data": user})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var params models.User
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := uh.repo.CreateUser(&params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Created user successfully", "data": params})
}
