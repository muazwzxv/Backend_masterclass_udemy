package users

import (
	"time"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	CreatedAt *time.Time
	DeletedAt *time.Time
}

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
}

type UpdateUser struct {
  // TODO:
}

type UpdatePassword struct {
  // TODO:
}
