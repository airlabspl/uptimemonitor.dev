-- name: CreateMonitor :exec
INSERT INTO monitors (uuid, url, user_id)
VALUES (?, ? , ?);

-- name: ListMonitors :many
SELECT * FROM monitors
WHERE user_id = ?
ORDER BY id DESC;
