-- name: GetSessionByUUID :one
SELECT sessions.*, sqlc.embed(users) FROM sessions 
LEFT JOIN users ON sessions.user_id = users.id
WHERE sessions.uuid = ?;

-- name: CreateSession :one
INSERT INTO sessions (uuid, user_id, expires_at)
VALUES (?, ?, ?)
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE uuid = ?;