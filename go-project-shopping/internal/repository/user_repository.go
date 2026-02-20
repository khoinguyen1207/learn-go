package repository

import (
	"context"
	"fmt"
	"project-shopping/internal/db"
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

func (ur *userRepository) GetAllV2(ctx context.Context, search, orderBy, sort string, limit, offset int32) ([]sqlc.User, error) {
	allowedOrderBy := map[string]bool{
		"created_at": true,
		"id":         true,
	}
	if !allowedOrderBy[orderBy] {
		orderBy = "created_at"
	}

	direction := "DESC"
	if sort == "asc" {
		direction = "ASC"
	}

	query := fmt.Sprintf(`
		SELECT * FROM users 
		WHERE deleted_at IS NULL
		AND (
			$1::TEXT IS NULL 
			OR $1::TEXT = ''
			OR email ILIKE '%%' || $1::TEXT || '%%'
			OR fullname ILIKE '%%' || $1::TEXT || '%%'
		)
		ORDER BY %s %s 
		LIMIT $2 OFFSET $3;`, orderBy, direction)

	rows, err := db.GetDBPool().Query(ctx, query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []sqlc.User{}
	for rows.Next() {
		var i sqlc.User
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Email,
			&i.Password,
			&i.Fullname,
			&i.Age,
			&i.Status,
			&i.Level,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
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

func (ur *userRepository) FindByUUID(ctx context.Context, id uuid.UUID) (sqlc.User, error) {
	user, err := ur.db.FindUserByUUID(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
