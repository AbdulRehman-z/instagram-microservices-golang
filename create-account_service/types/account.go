package types

type (
	CreateAccountReqParams struct {
		Username string `json:"username" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Bio      string `json:"bio" validate:"required,min=3"`
		Age      int32  `json:"age" validate:"required,min=1"`
		Avatar   string `json:"avatar" validate:"required,min=3"`
		Status   string `json:"status" validate:"required,min=3"`
	}

	CreateAccountResParams struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Age      int32  `json:"age"`
		Avatar   string `json:"avatar"`
		Status   string `json:"status"`
	}

	UpdateAccountReqParams struct {
		Username string `json:"username" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Bio      string `json:"bio" validate:"required,min=3"`
		Age      int32  `json:"age" validate:"required,min=1"`
		Avatar   string `json:"avatar" validate:"required,min=3"`
		Status   string `json:"status" validate:"required,min=3"`
	}

	UpdateAccountResParams struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Age      int32  `json:"age"`
		Avatar   string `json:"avatar"`
		Status   string `json:"status"`
	}

	DeleteAccountReqParams struct {
		UniqueId string `json:"unique_id" validate:"required"`
	}

	DeleteAccountResParams struct {
		UniqueId string `json:"unique_id"`
	}

	GetAccountReqParams struct {
		UniqueId string `json:"unique_id" validate:"required"`
	}

	GetAccountResParams struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Age      int32  `json:"age"`
		Avatar   string `json:"avatar"`
		Status   string `json:"status"`
	}
)
