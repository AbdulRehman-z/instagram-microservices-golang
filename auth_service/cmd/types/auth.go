package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	RegisterUserReqParams struct {
		Email    string `json:"username" validate:"required,alpha,min=1,max=20"`
		Password string `json:"password" validate:"required,min=8"`
	}

	RegisterUserResParams struct {
		Email             string    `json:"username"`
		HashedPassword    string    `json:"hashed_password"`
		PasswordChangedAt time.Time `json:"password_changed_at"`
		CreatedAt         time.Time `json:"password_created_at"`
	}

	LoginUserReqParams struct {
		Email    string `json:"username" validate:"required,alpha,min=1,max=20"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginUserResParams struct {
		SessionId             uuid.UUID `json:"session_id"`
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiresAt string    `json:"refresh_token_expires_at"`
	}
)
