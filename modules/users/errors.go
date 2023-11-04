package users

import "errors"

var (
	ErrNotFound             = errors.New("user not found")
	ErrFailedToQueryUser    = errors.New("failed to query user")
	ErrFaileToGenerateToken = errors.New("failed to generate token")
)
