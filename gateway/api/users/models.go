package users

import (
	"time"

	usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type User struct {
	ID        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	*usersModule.LoginResponse
}
