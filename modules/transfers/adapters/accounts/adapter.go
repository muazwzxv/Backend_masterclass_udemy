package adapter

import (
	"context"
)

type IAccounts interface {
	ValidateAccount(ctx context.Context, accountID int64, currency string) (bool, error)
}

type AccountsAdapter struct {
	acccountsModule IAccounts
}

func NewAccountsAdapter(accounts IAccounts) *AccountsAdapter {
	return &AccountsAdapter{
		acccountsModule: accounts,
	}
}
