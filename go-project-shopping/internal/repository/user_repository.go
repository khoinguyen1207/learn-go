package repository

import (
	"context"
	"project-shopping/internal/db/sqlc"

	"github.com/google/uuid"
)

type userRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetAll(ctx context.Context, search, orderBy, sort string, limit, offset int32) ([]sqlc.User, error) {
	var (
		data []sqlc.User
		err  error
	)

	switch {
	case orderBy == "created_at" && sort == "asc":
		data, err = ur.db.ListUsersOrderByCreatedAtAsc(ctx, sqlc.ListUsersOrderByCreatedAtAscParams{
			Search: search,
			Limit:  limit,
			Offset: offset,
		})
	case orderBy == "created_at" && sort == "desc":
		data, err = ur.db.ListUsersOrderByCreatedAtDesc(ctx, sqlc.ListUsersOrderByCreatedAtDescParams{
			Search: search,
			Limit:  limit,
			Offset: offset,
		})
	case orderBy == "id" && sort == "asc":
		data, err = ur.db.ListUsersOrderByIdAsc(ctx, sqlc.ListUsersOrderByIdAscParams{
			Search: search,
			Limit:  limit,
			Offset: offset,
		})
	case orderBy == "id" && sort == "desc":
		data, err = ur.db.ListUsersOrderByIdDesc(ctx, sqlc.ListUsersOrderByIdDescParams{
			Search: search,
			Limit:  limit,
			Offset: offset,
		})
	}

	if err != nil {
		return []sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) CountUsers(ctx context.Context, search string) (int64, error) {
	count, err := ur.db.CountUsers(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *userRepository) Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	data, err := ur.db.CreateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	data, err := ur.db.UpdateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Delete(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	data, err := ur.db.SoftDeleteUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) Restore(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	data, err := ur.db.RestoreUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return data, nil
}

func (ur *userRepository) FindById(id string) bool {
	return false
}

func (ur *userRepository) FindByEmail(email string) bool {
	return false
}
