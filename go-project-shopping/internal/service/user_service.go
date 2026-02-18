package service

import (
	"context"
	"database/sql"
	"errors"
	"project-shopping/internal/config"
	"project-shopping/internal/db/sqlc"
	"project-shopping/internal/repository"
	"project-shopping/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetUsers(ctx context.Context, search string, orderBy, sort string, page, limit int32) ([]sqlc.User, int32, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = config.DEFAULT_LIMIT
	}

	offset := (page - 1) * limit

	if orderBy == "" {
		orderBy = "created_at"
	}

	if sort == "" {
		sort = "desc"
	}

	users, err := us.repo.GetAllV2(ctx, search, orderBy, sort, limit, offset)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError(err, "Failed to get users", utils.CodeBadRequest)
	}

	totalUsers, err := us.repo.CountUsers(ctx, search)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError(err, "Failed to get total users", utils.CodeBadRequest)
	}

	return users, int32(totalUsers), nil
}

func (us *userService) GetUserByID(id string) error {
	return nil
}

func (us *userService) CreateUser(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error) {
	input.Email = utils.NormalizeString(input.Email)

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return sqlc.User{}, utils.WrapError(err, "Failed to hash password", utils.CodeBadRequest)
	}
	input.Password = string(hashedPassword)

	user, err := us.repo.Create(ctx, input)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return sqlc.User{}, utils.NewError("Email already exists", utils.CodeConflict)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to create user", utils.CodeBadRequest)
	}

	return user, nil
}

func (us *userService) UpdateUser(ctx context.Context, input sqlc.UpdateUserParams) (sqlc.User, error) {
	if input.Password != nil {
		hashedPassword, err := utils.HashPassword(*input.Password)
		if err != nil {
			return sqlc.User{}, utils.WrapError(err, "Failed to hash password", utils.CodeBadRequest)
		}

		hashedPasswordStr := string(hashedPassword)
		input.Password = &hashedPasswordStr
	}

	user, err := us.repo.Update(ctx, input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found", utils.CodeNotFound)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to update user", utils.CodeBadRequest)
	}

	return user, nil
}

func (us *userService) DeleteUser(ctx context.Context, id string) (sqlc.User, error) {
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, utils.WrapError(err, "Invalid user ID format", utils.CodeBadRequest)
	}

	user, err := us.repo.Delete(ctx, uuidParsed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found or already deleted", utils.CodeNotFound)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to delete user", utils.CodeBadRequest)
	}

	return user, nil
}

func (us *userService) RestoreUser(ctx context.Context, id string) (sqlc.User, error) {
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		return sqlc.User{}, utils.WrapError(err, "Invalid user ID format", utils.CodeBadRequest)
	}

	user, err := us.repo.Restore(ctx, uuidParsed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found or not deleted", utils.CodeNotFound)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to restore user", utils.CodeBadRequest)
	}

	return user, nil
}
