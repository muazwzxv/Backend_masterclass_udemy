package accounts

import (
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (m *Module) CreateAccount(ctx context.Context, data *CreateAccount) (*db.Account, error) {
  acc, err := m.db.CreateAccount(ctx, db.CreateAccountParams{
    OwnerID: data.OwnerID,
    Balance: data.Balance,
    Currency: data.Currency,
  })
  if err != nil {
    m.log.Info(err, "CreateAccount")
    return nil, errors.Wrapf(err, "m.db.CreateAccount")
  }

  return &acc, nil
}

func (m *Module) FindAccount(ctx context.Context, id int64) (*db.Account, error) {
  acc, err := m.db.GetAccount(ctx, id)
  if err != nil {
    return nil, errors.Wrapf(err, "m.db.GetAccount")
  }
	return &acc, nil
}
