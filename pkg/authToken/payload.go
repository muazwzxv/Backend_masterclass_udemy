package authToken

import (
  "time"
  "github.com/google/uuid"

)

type Payload struct {
	ID        uuid.UUID
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
  tokenId, err := uuid.NewRandom()
  if err != nil {
    return nil, err
  }

  payload := &Payload{
    ID: tokenId,
    Username: username,
    IssuedAt: time.Now(),
    ExpiredAt: time.Now().Add(duration),
  }

  return payload, nil
}

func (p *Payload) Valid() error {
  if time.Now().After(p.ExpiredAt) {
    return ErrTokenExpired
  }
  return nil
}
