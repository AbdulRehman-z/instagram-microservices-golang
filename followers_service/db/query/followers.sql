-- name: FollowUser :one
INSERT INTO followers (leader_unique_id, follower_unique_id)
VALUES ($1, $2)
ON CONFLICT (leader_unique_id, follower_unique_id) DO NOTHING
RETURNING *;

-- name: UnfollowUser :one
DELETE FROM followers
WHERE leader_unique_id = $1 AND follower_unique_id = $2
RETURNING *;

-- name: GetFollowersCount :one
SELECT COUNT(*) FROM followers
WHERE leader_unique_id = $1;

-- name: GetFollowingCount :one
SELECT COUNT(*) FROM followers
WHERE follower_unique_id = $1;

-- name: GetFollowers :many
SELECT "follower_unique_id" FROM followers
WHERE leader_unique_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetFollowing :many
SELECT "leader_unique_id" FROM followers
WHERE follower_unique_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;