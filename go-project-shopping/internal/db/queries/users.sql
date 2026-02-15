-- name: CreateUser :one
INSERT INTO users (email, password, fullname, age, status, level) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    password = COALESCE(sqlc.narg(password), password),
    fullname = COALESCE(sqlc.narg(fullname), fullname),
    age = COALESCE(sqlc.narg(age), age),
    status = COALESCE(sqlc.narg(status), status),
    level = COALESCE(sqlc.narg(level), level)
WHERE id = sqlc.narg(id) AND deleted_at IS NULL
RETURNING *;