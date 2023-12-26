-- name: CreatePost :one
INSERT INTO "posts" (
    "unique_id",
    "url",
    "caption",
    "lat",
    "lng"
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetPostsByUniqueId :many
SELECT 
    "id",
    "unique_id",
    "created_at",
    "url",
    "caption",
    "lat",
    "lng"
FROM "posts" WHERE "unique_id" = $1;

-- name: DeletePostsByUniqueId :exec
DELETE
FROM "posts"
WHERE "unique_id" = $1;

-- name: GetPost :one
SELECT 
    "id",
    "unique_id",
    "created_at",
    "url",
    "caption",
    "lat",
    "lng"
FROM "posts" WHERE "id" = $1;

-- name: GetPosts :many
SELECT 
    "id",
    "unique_id",
    "created_at",
    "url",
    "caption",
    "lat",
    "lng"
FROM "posts" ORDER BY "created_at" DESC LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE "posts" SET
"url" = coalesce($2, "url"),
"caption" = coalesce($3, "caption"),
"lat" = coalesce($4, "lat"),
"lng" = coalesce($5, "lng")
WHERE "id" = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM "posts" WHERE "id" = $1;
