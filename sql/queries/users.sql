-- name: ListUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
  AND deleted_at IS NULL;
-- name: CreateUser :exec
INSERT INTO users (name)
VALUES ($1);
-- name: UpdateUser :exec
UPDATE users
SET name = $2
WHERE id = $1
  AND deleted_at IS NULL;
-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;