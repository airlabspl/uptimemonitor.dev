-- name: CreatePasswordReset :exec
INSERT INTO password_resets(user_id, token, created_at, expires_at)
VALUES (?, ?, ?, ?);

