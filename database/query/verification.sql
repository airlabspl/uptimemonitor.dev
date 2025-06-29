-- name: CreateVerification :exec
INSERT INTO verifications (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetVerificationByToken :one
SELECT verifications.*, sqlc.embed(users) 
FROM verifications 
JOIN users ON verifications.user_id = users.id 
WHERE token = ?
LIMIT 1;

-- name: VerifyUser :exec
UPDATE users
SET email_verified_at = ?
WHERE id = ?;

-- name: DeleteVerification :exec
DELETE FROM verifications
WHERE id = ?;

-- name: GetLatestUserVerification :one
SELECT * FROM verifications
WHERE user_id = ?
ORDER BY id DESC
LIMIT 1;