package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PaestoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaestoMaker(symmetricKey string) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d", chacha20poly1305.KeySize)
	}

	return &PaestoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (p *PaestoMaker) CreateToken(email string, uniqueId uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, uniqueId, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func (p *PaestoMaker) VerifyToken(token string) (*Payload, error) {

	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if !payload.Valid() {
		return nil, fmt.Errorf("token is expired")
	}

	return payload, nil
}
