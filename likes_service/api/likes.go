package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/likes_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/likes_service/util"
	"github.com/gofiber/fiber/v2"
)

// LikePost adds a like to the post
func (s *Server) LikePost(c *fiber.Ctx) error {
	var req types.LikePost
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.LikePostParams{
		PostID:       req.PostID,
		UserUniqueID: req.User_unique_id,
	}

	res, err := s.store.LikePost(c.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"post_id":        res.PostID,
		"user_unique_id": res.UserUniqueID,
	})
}

// UnlikePost removes a like from the post
func (s *Server) UnlikePost(c *fiber.Ctx) error {
	var req types.UnlikePost
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.UnlikePostParams{
		PostID:       req.PostID,
		UserUniqueID: req.User_unique_id,
	}

	res, err := s.store.UnlikePost(c.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"post_id":        res.PostID,
		"user_unique_id": res.UserUniqueID,
	})
}

// GetPostLikesCount returns the number of likes for the given post
func (s *Server) GetPostLikesCount(c *fiber.Ctx) error {
	var req types.GetPostLikesCount
	if err := c.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := s.store.GetPostLikesCount(c.Context(), req.PostID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"likes_count": res,
	})
}

// GetPostLikes returns the list of users who liked the given post
func (s *Server) GetPostLikes(c *fiber.Ctx) error {
	var req types.GetPostLikes
	if err := c.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.GetPostLikesParams{
		PostID: req.PostID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	res, err := s.store.GetPostLikes(c.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"likes": res,
	})
}

// LikeComment adds a like to the comment
func (s *Server) LikeComment(c *fiber.Ctx) error {
	var req types.LikeComments
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.LikeCommentParams{
		CommentID:    req.CommentID,
		UserUniqueID: req.User_unique_id,
	}

	res, err := s.store.LikeComment(c.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"comment_id":     res.CommentID,
		"user_unique_id": res.UserUniqueID,
	})
}

// UnlikeComment removes a like from the comment
func (s *Server) UnlikeComment(c *fiber.Ctx) error {
	var req types.UnlikeComments
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.UnlikeCommentParams{
		CommentID:    req.CommentID,
		UserUniqueID: req.User_unique_id,
	}

	res, err := s.store.UnlikeComment(c.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"comment_id":     res.CommentID,
		"user_unique_id": res.UserUniqueID,
	})
}

// GetCommentLikesCount returns the number of likes for the given comment
func (s *Server) GetCommentLikesCount(c *fiber.Ctx) error {
	var req types.GetCommentLikesCount
	if err := c.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := s.store.GetCommentLikesCount(c.Context(), req.CommentID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"likes_count": res,
	})
}

// HealthCheck returns the health of the service
func (s *Server) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "ok",
	})
}
