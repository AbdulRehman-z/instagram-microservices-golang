package types

import "github.com/google/uuid"

type (
	// FollowUserReqParams is the request params for the FollowUser handler
	FollowUserReqParams struct {
		LeaderUniqueId   uuid.UUID `query:"leader_unique_id" validate:"required"`
		FollowerUniqueId uuid.UUID `query:"follower_unique_id" validate:"required"`
	}

	// FollowUserResParams is the response params for the FollowUser handler
	FollowUserResParams struct {
		LeaderUniqueId   uuid.UUID `json:"leader_unique_id"`
		FollowerUniqueId uuid.UUID `json:"follower_unique_id"`
	}

	// UnfollowUserReqParams is the request params for the UnfollowUser handler
	UnfollowUserReqParams struct {
		LeaderUniqueId   uuid.UUID `query:"leader_unique_id" validate:"required"`
		FollowerUniqueId uuid.UUID `query:"follower_unique_id" validate:"required"`
	}

	// GetFollowersReqParams is the request params for the GetFollowers handler
	GetFollowersReqParams struct {
		UniqueId uuid.UUID `query:"unique_id" validate:"required"`
		Limit    int32     `query:"limit"`
		Offset   int32     `query:"offset"`
	}

	// GetFollowersResParams is the response params for the GetFollowers handler
	GetFollowersResParams struct {
		Followers []uuid.UUID `json:"followers"`
	}

	// GetFollowingsReqParams is the request params for the GetFollowings handler
	GetFollowingsReqParams struct {
		UniqueId uuid.UUID `query:"unique_id" validate:"required"`
	}

	// GetFollowingsResParams is the response params for the GetFollowings handler
	GetFollowingsResParams struct {
		Followings []uuid.UUID `json:"followings"`
	}

	// GetFollowersCountReqParams is the request params for the GetFollowersCount handler
	GetFollowersCountReqParams struct {
		UniqueId uuid.UUID `query:"unique_id" validate:"required"`
	}

	// GetFollowersCountResParams is the response params for the GetFollowersCount handler
	GetFollowersCountResParams struct {
		Count int `json:"count"`
	}

	// GetFollowingsCountReqParams is the request params for the GetFollowingsCount handler
	GetFollowingsCountReqParams struct {
		UniqueId uuid.UUID `query:"unique_id" validate:"required"`
	}

	// GetFollowingsCountResParams is the response params for the GetFollowingsCount handler
	GetFollowingsCountResParams struct {
		Count int `json:"count"`
	}
)
