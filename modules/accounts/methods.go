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
		m.log.Info(err, "m.db.CreateAccount")
		return nil, errors.Wrapf(err, "m.CreateAccount")
	}

	return convertToModuleAccount(acc), nil
}

func (m *Module) FindAccount(ctx context.Context, id int64) (*Account, error) {
	acc, err := m.db.GetAccount(ctx, id)
	if err != nil {
		m.log.Info(err, "m.db.GetAccount")
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return convertToModuleAccount(acc), nil
}

func (m *Module) ListAccounts(ctx context.Context, query *GetAccounts) ([]*Account, error) {
	accs, err := m.db.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  query.Limit,
		Offset: query.Offset,
	})
  if err != nil {
    m.log.Info(err, "m.db.ListAccounts")
    return nil, errors.Wrapf(err, "m.GetAccounts")
  }

  return convertToModuleAccountList(accs), nil
}
