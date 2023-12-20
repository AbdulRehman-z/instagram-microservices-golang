package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	RegisterUserReqParams struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterUserResParams struct {
		Email             string    `json:"email"`
		PasswordCreatedAt time.Time `json:"password_created_at"`
		PasswordChangedAt time.Time `json:"password_changed_at"`
	}

	LoginUserReqParams struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginUserResParams struct {
		SessionId             uuid.UUID `json:"session_id"`
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}

	ChangePasswordReqParams struct {
		Email       string `json:"email" validate:"required,email"`
		NewPassword string `json:"new_hashed_password" validate:"required,min=8"`
	}

	ChangePasswordResParams struct {
		Email             string    `json:"email"`
		PasswordChangedAt time.Time `json:"password_changed_at"`
	}
)
