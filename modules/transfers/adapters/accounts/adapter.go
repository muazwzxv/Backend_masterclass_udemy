package adapter

import (
	"context"

	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
)

type IAccounts interface {
	ValidateAccount(ctx context.Context, accountID int64, currency string) (bool, error)
}

type AccountsAdapter struct {
	acccountsModule IAccounts
}

func NewAccountsAdapter(accounts *accounts.Module) *AccountsAdapter {
	return &AccountsAdapter{
		acccountsModule: accounts,
	}
}
