-- name: CreateMonitor :exec
INSERT INTO monitors (uuid, url, user_id)
VALUES (?, ? , ?);

-- name: GetMonitors :many
SELECT * FROM monitors
WHERE user_id = ?;