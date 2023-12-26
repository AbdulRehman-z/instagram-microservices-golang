package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// CreatePostReqParams represents a request to create a new post.
	CreatePostReqParams struct {
		UniqueID string  `json:"unique_id" validate:"required"`
		Url      string  `json:"url" validate:"required"`
		Caption  string  `json:"caption" validate:"required"`
		Lat      float64 `json:"lat" validate:"required"`
		Lng      float64 `json:"lng" validate:"required"`
	}

	// CreatePostResParams represents a response for creating a new post.
	CreatePostResParams struct {
		Message string `json:"message"`
	}

	GetPostsReqParams struct {
		Limit  int32 `query:"limit" validate:"required,min=1,max=10"`
		Offset int32 `query:"limit" validate:"required,min=0,max=100"`
	}

	GetPostsResParams struct {
		Posts []struct {
			Id       int32   `json:"id"`
			UniqueID string  `json:"unique_id"`
			Url      string  `json:"url"`
			Caption  string  `json:"caption"`
			Lat      float64 `json:"lat"`
			Lng      float64 `json:"lng"`
		} `json:"posts"`
	}

	// CreatePostResParams represents a request to
	// fetch all posts belongs to the user having the unique_id.
	GetPostsByUniqueIDReqParams struct {
		UniqueID uuid.UUID `query:"unique_id" validate:"required"`
	}

	// GetPostsByUniqueIDResParams representsa a response of fetch all posts belongs to a user.
	GetPostsByUniqueIDResParams struct {
		Posts []struct {
			Id       int32   `json:"id"`
			UniqueID string  `json:"unique_id"`
			Url      string  `json:"url"`
			Caption  string  `json:"caption"`
			Lat      float64 `json:"lat"`
			Lng      float64 `json:"lng"`
		} `json:"posts"`
	}

	// GetPostByidReqParams represents a request that fetch post by id.
	GetPostByidReqParams struct {
		Id int32 `query:"id" validate:"required"`
	}

	// GetPostByidResParams represents a response that fetch post by id.
	GetPostByidResParams struct {
		Id        int32     `json:"id"`
		UniqueID  string    `json:"unique_id"`
		Url       string    `json:"url"`
		Caption   string    `json:"caption"`
		Lat       float64   `json:"lat"`
		Lng       float64   `json:"lng"`
		CreatedAt time.Time `json:"created_at"`
	}

	// UpdatePostReqParams represents a request to update a post.
	UpdatePostReqParams struct {
		Id       int32   `json:"id" validate:"required"`
		UniqueID string  `json:"unique_id" validate:"required"`
		Url      string  `json:"url" validate:"required"`
		Caption  string  `json:"caption" validate:"required"`
		Lat      float64 `json:"lat" validate:"required"`
		Lng      float64 `json:"lng" validate:"required"`
	}

	// UpdatePostResParams represents a response for updating a post.
	UpdatePostResParams struct {
		Message string `json:"message"`
	}

	// DeletePostReqParams represents a request to delete a post.
	DeletePostReqParams struct {
		Id int32 `query:"id" validate:"required"`
	}

	// DeletePostResParams represents a response for deleting a post.
	DeletePostResParams struct {
		Message string `json:"message"`
	}

	DeletePostsByUniqueIDReqParams struct {
		UniqueId uuid.UUID `query:"unique_id" validate:"required"`
	}

	DeleteAccountByUniqueIDResParams struct {
		Message string `json:"message"`
	}
)
