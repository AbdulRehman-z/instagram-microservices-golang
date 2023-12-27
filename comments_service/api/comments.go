package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/comments_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/comments_service/util"
	"github.com/gofiber/fiber/v2"
)

// CreateComment creates a new comment
func (s *Server) CreateComment(ctx *fiber.Ctx) error {
	var req types.CreateCommentReqParams
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payload := ctx.Locals(authorizationPayloadKey).(*token.Payload)

	args := db.CreateCommentParams{
		PostID:       req.Post_id,
		UserUniqueID: payload.UniqueId,
		Content:      req.Content,
	}

	comment, err := s.store.CreateComment(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": comment,
	})
}

// GetCommentsByPostID returns the list of comments for the given post
func (s *Server) GetComments(ctx *fiber.Ctx) error {
	var req types.GetCommentsReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.GetCommentsParams{
		PostID: req.Post_id,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	comments, err := s.store.GetComments(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": comments,
	})
}

func (s *Server) UpdateComment(ctx *fiber.Ctx) error {
	var req types.UpdateCommentReqParams
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.UpdateCommentParams{
		ID:      req.Comment_id,
		Content: req.Content,
	}

	comment, err := s.store.UpdateComment(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": comment,
	})
}

func (s *Server) GetCommentsCount(ctx *fiber.Ctx) error {
	var req types.GetCommentsCountReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	count, err := s.store.GetCommentsCount(ctx.Context(), req.Post_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": count,
	})
}

func (s *Server) DeleteComment(ctx *fiber.Ctx) error {
	var req types.DeleteCommentReqParams
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := s.store.DeleteComment(ctx.Context(), req.Comment_id); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": "comment deleted",
	})
}
