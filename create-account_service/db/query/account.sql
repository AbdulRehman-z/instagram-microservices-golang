-- name: CreateAccount :one
INSERT INTO "accounts" (
  "username",
  "avatar",
  "age",
  "bio",
  "status"
)
VALUES(
    $1,$2,$3,$4,$5
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
WHERE "id" = $1
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM "accounts"
WHERE "id" = $1;
