package accounts

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

type CreateAccount struct {
	OwnerID  int64       `json:"owner_id" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required,oneof=USD EUR"`
}

type GetAccountrequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type GetAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type Account struct {
	ID       int64       `json:"id"`
	OwnerID  int64       `json:"owner_id"`
	Balance  int64       `json:"balance"`
	Currency db.Currency `json:"currency"`
}
