-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES (?, ?, ?)
RETURNING *;