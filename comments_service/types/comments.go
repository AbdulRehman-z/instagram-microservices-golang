package types

import "github.com/google/uuid"

type (
	CreateCommentReqParams struct {
		Post_id        int32     `json:"post_id" validate:"required"`
		Unique_User_id uuid.UUID `json:"unique_user_id" validate:"required"`
		Content        string    `json:"content" validate:"required"`
	}

	GetCommentsCountReqParams struct {
		Post_id int32 `json:"post_id" validate:"required"`
	}

	GetCommentsReqParams struct {
		Post_id int32 `query:"post_id" validate:"required"`
		Offset  int32 `query:"offset" validate:"required"`
		Limit   int32 `query:"limit" validate:"required"`
	}

	UpdateCommentReqParams struct {
		Comment_id int32  `json:"comment_id" validate:"required"`
		Content    string `json:"content" validate:"required"`
	}

	DeleteCommentReqParams struct {
		Comment_id int32 `json:"comment_id" validate:"required"`
	}
)
