-- name: LikePost :one
INSERT INTO post_likes (post_id, user_unique_id)
VALUES ($1, $2)
ON CONFLICT (post_id, user_unique_id) DO NOTHING
RETURNING *;

-- name: UnlikePost :one
DELETE FROM post_likes
WHERE post_id = $1 AND user_unique_id = $2
RETURNING *;

-- name: LikeComment :one
INSERT INTO comment_likes (comment_id, user_unique_id)
VALUES ($1, $2)
ON CONFLICT (comment_id, user_unique_id) DO NOTHING
RETURNING *;

-- name: UnlikeComment :one
DELETE FROM comment_likes
WHERE comment_id = $1 AND user_unique_id = $2
RETURNING *;

-- name: GetPostLikesCount :one
SELECT COUNT(*) FROM post_likes
WHERE post_id = $1;

-- name: GetPostLikes :many
SELECT "user_unique_id" FROM post_likes
WHERE post_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetCommentLikesCount :one
SELECT COUNT(*) FROM comment_likes
WHERE comment_id = $1;

-- name: GetCommentLikes :many
SELECT "user_unique_id" FROM comment_likes
WHERE comment_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

