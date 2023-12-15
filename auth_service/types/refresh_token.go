package types

import (
	"time"
)

type (
	RenewAccessTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RenewAccessTokenResponse struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}
)
