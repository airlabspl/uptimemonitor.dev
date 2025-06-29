-- name: CreateVerification :exec
INSERT INTO verifications (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetVerificationByToken :one
SELECT verifications.*, sqlc.embed(users) 
FROM verifications 
JOIN users ON verifications.user_id = users.id 
WHERE token = ?;

-- name: VerifyUser :exec
UPDATE users
SET email_verified_at = ?
WHERE id = ?;

-- name: DeleteVerification :exec
DELETE FROM verifications
WHERE id = ?;