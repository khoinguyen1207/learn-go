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

-- name: SoftDeleteUser :one
UPDATE users
SET deleted_at = NOW()
WHERE uuid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: RestoreUser :one
UPDATE users
SET deleted_at = NULL
WHERE uuid = $1 AND deleted_at IS NOT NULL
RETURNING *;

-- name: ListUsersOrderByCreatedAtAsc :many
SELECT * FROM users 
WHERE deleted_at IS NULL
AND (
    sqlc.narg(search)::TEXT IS NULL 
    OR sqlc.narg(search)::TEXT = ''
    OR email ILIKE '%' || sqlc.arg(search) || '%'
    OR fullname ILIKE '%' || sqlc.arg(search) || '%'
)
ORDER BY created_at ASC
LIMIT $1 OFFSET $2;

-- name: ListUsersOrderByCreatedAtDesc :many
SELECT * FROM users 
WHERE deleted_at IS NULL
AND (
    sqlc.narg(search)::TEXT IS NULL 
    OR sqlc.narg(search)::TEXT = ''
    OR email ILIKE '%' || sqlc.arg(search) || '%'
    OR fullname ILIKE '%' || sqlc.arg(search) || '%'
)
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListUsersOrderByIdAsc :many
SELECT * FROM users 
WHERE deleted_at IS NULL
AND (
    sqlc.narg(search)::TEXT IS NULL 
    OR sqlc.narg(search)::TEXT = ''
    OR email ILIKE '%' || sqlc.arg(search) || '%'
    OR fullname ILIKE '%' || sqlc.arg(search) || '%'
)
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: ListUsersOrderByIdDesc :many
SELECT * FROM users 
WHERE deleted_at IS NULL
AND (
    sqlc.narg(search)::TEXT IS NULL 
    OR sqlc.narg(search)::TEXT = ''
    OR email ILIKE '%' || sqlc.arg(search) || '%'
    OR fullname ILIKE '%' || sqlc.arg(search) || '%'
)
ORDER BY id DESC
LIMIT $1 OFFSET $2;