package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/followers_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/followers_service/types"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) FollowUser(ctx *fiber.Ctx) error {
	var req types.FollowUserReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.FollowUserParams{
		LeaderUniqueID:   req.LeaderUniqueId,
		FollowerUniqueID: req.FollowerUniqueId,
	}

	res, err := s.store.FollowUser(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"leader_unique_id":   res.LeaderUniqueID,
		"follower_unique_id": res.FollowerUniqueID,
	})
}

// UnfollowUser removes the follower from the leader's followers list
func (s *Server) UnfollowUser(ctx *fiber.Ctx) error {
	var req types.UnfollowUserReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.UnfollowUserParams{
		LeaderUniqueID:   req.LeaderUniqueId,
		FollowerUniqueID: req.FollowerUniqueId,
	}

	res, err := s.store.UnfollowUser(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"leader_unique_id":   res.LeaderUniqueID,
		"follower_unique_id": res.FollowerUniqueID,
	})
}

// GetFollowers returns the list of followers for the given user
func (s *Server) GetFollowers(ctx *fiber.Ctx) error {
	var req types.GetFollowersReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.GetFollowersParams{
		LeaderUniqueID: req.UniqueId,
		Limit:          req.Limit,
		Offset:         req.Offset,
	}

	res, err := s.store.GetFollowers(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"followers": res,
	})
}

// GetFollowings returns the list of followings for the given user
func (s *Server) GetFollowings(ctx *fiber.Ctx) error {
	var req types.GetFollowingsReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.GetFollowingParams{
		FollowerUniqueID: req.UniqueId,
	}

	res, err := s.store.GetFollowing(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"followings": res,
	})
}

// HealthCheck returns the health of the service
func (s *Server) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

//	!------- NOTE: The following two handlers will be mostly called by the stream service ------!
//
// GetFollowersCount returns the count of followers for the given user
func (s *Server) GetFollowersCount(ctx *fiber.Ctx) error {
	var req types.GetFollowersCountReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	count, err := s.store.GetFollowersCount(ctx.Context(), req.UniqueId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
	})
}

// GetFollowingsCount returns the count of followings for the given user
func (s *Server) GetFollowingsCount(ctx *fiber.Ctx) error {
	var req types.GetFollowingsCountReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	count, err := s.store.GetFollowingCount(ctx.Context(), req.UniqueId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
	})
}
