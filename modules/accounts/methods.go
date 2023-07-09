package accounts

import (
	"database/sql"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (m *Module) CreateAccount(ctx context.Context, data *CreateAccount) (*Account, error) {
	acc, err := m.db.CreateAccount(ctx, db.CreateAccountParams{
		OwnerID:  data.OwnerID,
		Balance:  data.Balance,
		Currency: data.Currency,
	})
	if err != nil {
		m.log.Errorf("m.db.CreateAccount: %v", err)
		return nil, errors.Wrapf(err, "m.CreateAccount")
	}

	return convertToModuleAccount(acc), nil
}

func (m *Module) FindAccount(ctx context.Context, id int64) (*Account, error) {
	acc, err := m.db.GetAccount(ctx, id)
	if err != nil {
		m.log.Errorf("m.db.GetAccount: %v", err)
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrapf(err, "m.FindAccount")
	}
	return convertToModuleAccount(acc), nil
}

func (m *Module) ListAccounts(ctx context.Context, query *GetAccounts) ([]*Account, error) {
	accs, err := m.db.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  query.Limit,
		Offset: query.Offset,
	})
	if err != nil {
		m.log.Errorf("m.db.ListAccounts: %v", err)
		return nil, errors.Wrapf(err, "m.GetAccounts")
	}

	return convertToModuleAccountList(accs), nil
}

func (m *Module) ValidateAccount(ctx context.Context, accountID int64, currency string) (bool, error) {
	acc, err := m.db.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, ErrNotFound
		}
		return false, errors.Wrapf(err, "m.ValidateAccount")
	}

	if string(acc.Currency) != currency {
		return false, errors.Errorf("account [%d] currency mismatch: %s and %s", accountID, string(acc.Currency), currency)
	}

	return true, nil
}
