package accounts

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

type CreateAccount struct {
	OwnerID  int64
	Balance  int64
	Currency db.Currency
}
