package accounts

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

func convertToModuleAccount(acc db.Account) *Account {
	return &Account{
		ID:        acc.ID,
		OwnerID:   acc.OwnerID,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Time,
		DeletedAt: acc.DeletedAt.Time,
	}
}

func convertToModuleAccountList(accs []db.Account) []*Account {
	accounts := make([]*Account, len(accs))
	for _, acc := range accs {
		account := &Account{
			ID:        acc.ID,
			OwnerID:   acc.OwnerID,
			Balance:   acc.Balance,
			Currency:  acc.Currency,
			CreatedAt: acc.CreatedAt.Time,
			DeletedAt: acc.DeletedAt.Time,
		}

    accounts = append(accounts, account)
	}

  return accounts
}
