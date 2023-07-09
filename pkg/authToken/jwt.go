package authToken

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKey = 32

type Jwt struct {
	secretKey string
}

var _ IToken = (*Jwt)(nil)

func NewJwt(secretKey string) (*Jwt, error) {
	if len(secretKey) < minSecretKey {
		return nil, ErrInvalidKeySize
	}

	return &Jwt{
		secretKey: secretKey,
	}, nil
}

func (t *Jwt) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

  // payload must implements `Valid() error` function
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

  return jwtToken.SignedString([]byte(t.secretKey))
}

func (t *Jwt) VerifyToken(token string) (*Payload, error) {
  jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(jwtToken *jwt.Token) (interface{}, error) {
    // t.Method returns an interface, we try to cast to an implementation
    _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
    if !ok {
      // We're using HS256 and HMAC, if casting to HMAC signing method fails, it means the token is invalid or tempered with
      return nil, ErrInvalidToken
    }

    return []byte(t.secretKey), nil
  })

  if err != nil {
    jwtErr, ok := err.(*jwt.ValidationError)
    if ok && errors.Is(jwtErr.Inner, ErrTokenExpired) {
      return nil, ErrTokenExpired
    }
    return nil, ErrInvalidToken
  }

  payload, ok := jwtToken.Claims.(*Payload)

  if !ok {
    return nil, ErrInvalidToken
  }

  return payload, nil
}
