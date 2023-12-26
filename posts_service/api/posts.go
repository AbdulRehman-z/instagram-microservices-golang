package api

import (
	"database/sql"

	db "github.com/AbdulRehman-z/instagram-microservices/posts_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/posts_service/util"
	"github.com/gofiber/fiber/v2"
)

// CreatePostRequest represents a request to create a new post.
func (s *Server) CreatePost(ctx *fiber.Ctx) error {
	var req types.CreatePostReqParams
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payload := ctx.Locals(authorizationPayloadKey).(*token.Payload)
	args := db.CreatePostParams{
		UniqueID: payload.UniqueId,
		Url:      req.Url,
		Caption:  sql.NullString{String: req.Caption, Valid: req.Caption != ""},
		Lat:      sql.NullFloat64{Float64: req.Lat, Valid: true},
		Lng:      sql.NullFloat64{Float64: req.Lng, Valid: true},
	}

	post, err := s.store.CreatePost(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response := types.CreatePostResParams{
		Message: string(post.ID),
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})
}

// GetPostById represents a request to get a post by id.
func (s *Server) GetPostById(ctx *fiber.Ctx) error {
	var req types.GetPostByidReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	post, err := s.store.GetPost(ctx.Context(), req.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response := types.GetPostByidResParams{
		Id:        post.ID,
		UniqueID:  post.UniqueID.String(),
		Url:       post.Url,
		Caption:   post.Caption.String,
		Lat:       post.Lat.Float64,
		Lng:       post.Lng.Float64,
		CreatedAt: post.CreatedAt.Time,
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})
}

// GetPostsByUniqueId represents a request to get posts by unique id.
func (s *Server) GetPostsByUniqueId(ctx *fiber.Ctx) error {
	var req types.GetPostsByUniqueIDReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	posts, err := s.store.GetPostsByUniqueId(ctx.Context(), req.UniqueID)
	if err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": posts,
	})
}

// UpdatePost represents a request to update a post.
func (s *Server) UpdatePost(ctx *fiber.Ctx) error {
	var req types.UpdatePostReqParams
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.UpdatePostParams{
		ID:      req.Id,
		Url:     req.Url,
		Caption: sql.NullString{String: req.Caption, Valid: req.Caption != ""},
		Lat:     sql.NullFloat64{Float64: req.Lat, Valid: true},
		Lng:     sql.NullFloat64{Float64: req.Lng, Valid: true},
	}

	post, err := s.store.UpdatePost(ctx.Context(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": post,
	})
}

// DeletePost represents a request to delete a post.
func (s *Server) DeletePost(ctx *fiber.Ctx) error {
	var req types.DeletePostReqParams

	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := s.store.DeletePost(ctx.UserContext(), req.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "posts deleted",
	})
}

// DeletePostsByUniqueId represents a request to delete posts by unique id.
func (s *Server) DeletePostsByUniqueId(ctx *fiber.Ctx) error {
	var req types.DeletePostsByUniqueIDReqParams
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := s.store.DeletePostsByUniqueId(ctx.Context(), req.UniqueId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "posts deleted",
	})
}

// Health represents a request to check the health of the server.
func (s *Server) Health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "ok",
	})
}
