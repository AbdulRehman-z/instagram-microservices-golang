-- name: RegisterUser :one
INSERT INTO "users" (
        "email",
        "hashed_password"
    )
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM "users"
WHERE "email" = $1 
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
    hashed_password = coalesce(sqlc.narg('hashed_password'), hashed_password),
    password_changed_at = coalesce(sqlc.narg('password_changed_at'), password_changed_at),
    email = coalesce(sqlc.narg('email'), email),
    is_email_verified = coalesce(sqlc.narg('is_email_verified'), is_email_verified)
WHERE username = sqlc.arg(username)
RETURNING *;    