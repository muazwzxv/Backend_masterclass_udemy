package authToken

import "fmt"

var (
	ErrInvalidKeySize = fmt.Errorf("Invalid key size: must be at least %d characters", minSecretKey)
	ErrTokenExpired   = fmt.Errorf("Token has expired")
	ErrInvalidToken   = fmt.Errorf("Token is invalid")
)
