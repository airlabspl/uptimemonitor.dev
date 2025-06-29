-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: CountAdminUsers :one
SELECT COUNT(*) AS count FROM users WHERE is_admin = TRUE;

-- name: CreateAdminUser :one
INSERT INTO users (name, email, password_hash, is_admin)
VALUES (?, ?, ?, TRUE)
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = ?
WHERE id = ?;