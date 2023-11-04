package accounts

import (
	"time"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

type Account struct {
	ID        int64       `json:"id"`
	OwnerID   int64       `json:"owner_id"`
	Balance   int64       `json:"balance"`
	Currency  db.Currency `json:"currency"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

type CreateAccount struct {
	OwnerID  int64
	Balance  int64
	Currency db.Currency
}

type GetAccounts struct {
	Limit  int32
	Offset int32
}

