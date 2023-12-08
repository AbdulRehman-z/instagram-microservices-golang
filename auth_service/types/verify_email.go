package types

type (
	VerifyEmailReqParams struct {
		EmailId    int64  `query:"email_id" validate:"required"`
		SecretCode string `query:"secret_code" validate:"required,gte=6"`
	}

	VerifyEmailResponse struct {
		IsEmailVerified bool `json:"is_email_verified"`
	}
)
