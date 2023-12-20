package token

import (
	"time"

	"github.com/google/uuid"
)

type TokenMaker interface {
	CreateToken(email string, uniqueId uuid.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
