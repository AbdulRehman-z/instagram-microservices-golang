-- name: FollowUser :one
INSERT INTO followers (leader_unique_id, follower_unique_id)
VALUES ($1, $2)
ON CONFLICT (leader_unique_id, follower_unique_id) DO NOTHING
RETURNING *;

-- name: UnfollowUser :one
DELETE FROM followers
WHERE leader_unique_id = $1 AND follower_unique_id = $2
RETURNING *;

-- name: GetFollowers :many
Explain select * from followers where leader_unique_id = $1;

-- name: GetFollowing :many
Explain select * from followers where follower_unique_id = $1;