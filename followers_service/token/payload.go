package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	UniqueId  uuid.UUID `json:"unique_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(email string, uniqueId uuid.UUID, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		Id:        id,
		Email:     email,
		UniqueId:  uniqueId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (payoad *Payload) Valid() bool {
	return time.Now().Unix() < payoad.ExpiredAt.Unix()
}
