-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  username, email, secret_code
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetVerifyEmail :one
SELECT 
    verify_emails.*, 
    users.is_user_verified
FROM verify_emails
LEFT JOIN users ON users.username = verify_emails.username
WHERE verify_emails.secret_code = $1;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
  is_used = TRUE
WHERE 
  id = sqlc.arg(id)
  AND secret_code = sqlc.arg(secret_code)
  AND is_used = FALSE
  AND expired_at > now()
RETURNING *;
