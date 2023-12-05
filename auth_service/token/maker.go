package token

import "time"

type TokenMaker interface {
	CreateToken(email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
