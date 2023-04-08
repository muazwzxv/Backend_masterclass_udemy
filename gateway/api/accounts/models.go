package accounts

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

type CreateAccount struct {
	OwnerID  int64       `json:"owner_id" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required,oneof=USD EUR"`
}

type Account struct {
	ID       int64       `json:"id"`
	OwnerID  int64       `json:"owner_id"`
	Balance  int64       `json:"balance"`
	Currency db.Currency `json:"currency"`
}
