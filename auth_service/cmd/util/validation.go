package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Message     string // Add a Message field to hold a custom error message
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate = validator.New(validator.WithRequiredStructEnabled())

var MyValidator = &XValidator{
	validator: validate,
}

func (v *XValidator) Validate(req interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(req)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Message = generateErrorMessage(err) // Generate custom error message
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func generateErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s should be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s should be at most %s characters long", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

func CheckValidationErrors(req interface{}) *fiber.Error {
	if errs := MyValidator.Validate(req); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, err.Message) // Use the custom error message
		}

		return fiber.NewError(fiber.StatusBadRequest, strings.Join(errMsgs, " and "))
	}
	return nil
}
