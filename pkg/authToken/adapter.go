package authToken

import "time"

type IToken interface {
  CreateToken(username string, duration time.Duration) (string, error)
  VerifyToken(token string) (*Payload, error)
} 
