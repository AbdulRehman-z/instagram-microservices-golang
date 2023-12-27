-- name: CreateComment :one
INSERT INTO comments (post_id, user_unique_id, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetComments :many
SELECT * FROM comments
WHERE post_id = $1
LIMIT $2
OFFSET $3; 

-- name: UpdateComment :one
UPDATE comments
SET content = $1
WHERE id = $2
RETURNING *;

-- name: GetCommentsCount :one
SELECT COUNT(*)
FROM comments
WHERE post_id = $1;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;
