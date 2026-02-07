package handlers

import (
	"errors"
	"go-sqlc/internal/db/sqlc"
	"go-sqlc/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	id := ctx.Param("uuid")
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID"})
		return
	}

	user, err := uh.repo.FindByUuid(ctx, parsedUuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Fetched user successfully", "data": user})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var params sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := uh.repo.CreateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Created user successfully", "data": user})
}
