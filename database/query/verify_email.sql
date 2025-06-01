-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  username, email, secret_code
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE username = $1 AND secret_code= $2;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
  is_used = sqlc.arg(is_used)
WHERE 
  id = sqlc.arg(id)
RETURNING *;
