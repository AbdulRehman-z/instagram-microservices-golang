package token

type TokenVerifier interface {
	VerifyToken(token string) (*Payload, error)
}
