package authToken

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type Paseto struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaseto(symmetricKey string) (IToken, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	pasetoToken := &Paseto{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return pasetoToken, nil
}

func (p *Paseto) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

func (p *Paseto) VerifyToken(token string) (*Payload, error) {
  payload := &Payload{}
  err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
  if err != nil {
    return nil, ErrInvalidToken
  }

  err = payload.Valid()
  if err != nil {
    return nil, err
  }

	return payload, nil
}
