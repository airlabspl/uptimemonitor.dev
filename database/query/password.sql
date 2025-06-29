-- name: CreatePasswordReset :exec
INSERT INTO password_resets(user_id, token, created_at, expires_at)
VALUES (?, ?, ?, ?);

-- name: GetPasswordResetByToken :one
SELECT password_resets.*, sqlc.embed(users)
FROM password_resets
LEFT JOIN users ON users.id = password_resets.user_id
WHERE token = ?;