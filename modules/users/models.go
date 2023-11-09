package users

import (
	"time"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	CreatedAt *time.Time
	DeletedAt *time.Time
}

type CreateUser struct {
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
}

type LoginUserRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	Token    string
	UserData User
}

type UpdateUser struct {
	// TODO:
}

type UpdatePassword struct {
	// TODO:
}
