package adapter

import (
	"context"
)

func (adpt *AccountsAdapter) ValidateAccount(ctx context.Context, accountID int64, currency string) (bool, error) {
	return adpt.acccountsModule.ValidateAccount(ctx, accountID, currency)
}
