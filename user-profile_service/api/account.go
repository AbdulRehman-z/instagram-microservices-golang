package api

import (
	"database/sql"
	"fmt"

	db "github.com/AbdulRehman-z/instagram-microservices/create-account_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/token"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/create-account_service/util"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

// CreateAccount creates a new account
func (s *Server) CreateAccount(c *fiber.Ctx) error {
	var req types.CreateAccountReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	args := db.CreateAccountParams{
		Email:    req.Email,
		Avatar:   req.Avatar,
		Username: req.Username,
		Age:      req.Age,
		Bio:      req.Bio,
		Status:   req.Status,
	}

	account, err := s.store.CreateAccount(c.UserContext(), args)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code.Name() == "unique_violation" {
				return fiber.NewError(fiber.StatusConflict, err.Error())
			}
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message":   "success",
		"unique_id": account.UniqueID,
	})
}

// GetAccount gets an account
func (s *Server) GetAccount(c *fiber.Ctx) error {
	payload := c.Locals(authorizationPayloadKey).(*token.Payload)
	fmt.Println("unique_id: ", payload.UniqueId)

	account, err := s.store.GetAccountByUniqueID(c.UserContext(), payload.UniqueId.String())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"account": account,
	})
}

// UpdateAccount updates an account
func (s *Server) UpdateAccount(c *fiber.Ctx) error {
	var req types.UpdateAccountReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payload := c.Locals(authorizationPayloadKey).(*token.Payload)
	if payload.UniqueId.String() != req.Username {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	args := db.UpdateAccountParams{
		UniqueID: payload.UniqueId.String(),
		Avatar:   sql.NullString{String: req.Avatar, Valid: req.Avatar != ""},
		Username: sql.NullString{String: req.Username, Valid: req.Username != ""},
		Age:      sql.NullInt32{Int32: req.Age, Valid: req.Age != 0},
		Bio:      sql.NullString{String: req.Bio, Valid: req.Bio != ""},
		Status:   sql.NullString{String: req.Status, Valid: req.Status != ""},
	}

	account, err := s.store.UpdateAccount(c.UserContext(), args)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"account": account,
	})
}

// DeleteAccount deletes an account
func (s *Server) DeleteAccount(c *fiber.Ctx) error {
	var req types.DeleteAccountReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payload := c.Locals(authorizationPayloadKey).(*token.Payload)
	if payload.UniqueId.String() != req.UniqueId {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	err := s.store.DeleteAccountByUniqueID(c.UserContext(), req.UniqueId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "success",
	})
}
