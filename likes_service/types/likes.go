package types

import "github.com/google/uuid"

type (
	LikePost struct {
		PostID         int32     `json:"post_id" validate:"required"`
		User_unique_id uuid.UUID `json:"user_unique_id" validate:"required"`
	}

	UnlikePost struct {
		PostID         int32     `json:"post_id" validate:"required"`
		User_unique_id uuid.UUID `json:"user_unique_id" validate:"required"`
	}

	GetPostLikes struct {
		PostID int32 `json:"post_id" validate:"required"`
		Limit  int32 `json:"limit" validate:"required"`
		Offset int32 `json:"offset" validate:"required"`
	}

	GetPostLikesCount struct {
		PostID int32 `json:"post_id" validate:"required"`
	}

	LikeComments struct {
		CommentID      int32     `json:"comment_id" validate:"required"`
		User_unique_id uuid.UUID `json:"user_unique_id" validate:"required"`
	}

	UnlikeComments struct {
		CommentID      int32     `json:"comment_id" validate:"required"`
		User_unique_id uuid.UUID `json:"user_unique_id" validate:"required"`
	}

	GetCommentLikes struct {
		CommentID int32 `json:"comment_id" validate:"required"`
	}

	GetCommentLikesCount struct {
		CommentID int32 `json:"comment_id" validate:"required"`
	}
)
