package accounts

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

func convertToAccountResponse(acc *db.Account) Account {
	return Account{
		ID:       acc.ID,
		OwnerID:  acc.OwnerID,
		Balance:  acc.Balance,
		Currency: acc.Currency,
	}
}
