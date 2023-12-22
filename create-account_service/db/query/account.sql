-- name: CreateAccount :one
INSERT INTO "accounts" (
  "email",
  "username",
  "avatar",
  "age",
  "bio",
  "status"
)
VALUES(
    $1,$2,$3,$4,$5,$6
) 
RETURNING *;

-- name: UpdateAccount :one
UPDATE "accounts"
SET
    "username" = coalesce(sqlc.narg('username'), username),
    "avatar" = coalesce(sqlc.narg('avatar'), avatar),
    "age" = coalesce(sqlc.narg('age'), age),
    "bio" = coalesce(sqlc.narg('bio'), bio),
    "status" = coalesce(sqlc.narg('status'), status)
WHERE "unique_id" = $1
RETURNING *;

-- name: GetAccountByUniqueID :one
SELECT *
FROM "accounts"
WHERE "unique_id" = $1;

-- name: DeleteAccountByUniqueID :exec
DELETE FROM "accounts"
WHERE "unique_id" = $1;
