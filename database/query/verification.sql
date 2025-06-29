-- name: CreateVerification :exec
INSERT INTO verifications (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetVerificationByToken :one
SELECT * FROM verifications WHERE token = ?;