-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;
