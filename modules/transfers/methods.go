package transfers

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/pkg/errors"
)

func (m *Module) TransferTransaction(ctx context.Context, req *TransferRequest) (*db.TransferTxResult, error) {

  // validate accounts and currency 
  isFromAccountValid, err := m.accounts.ValidateAccount(ctx, req.FromAccountID, req.Currency)
  if err != nil {
		m.log.Errorf("m.accounts.ValidateAccount: %v", err)
		return nil, errors.Wrapf(err, "m.TransferTransaction")
  }

  if !*isFromAccountValid {
    return nil, errors.Wrapf(err, "m.TransferTransaction")
  }

  isToAccountValid, err := m.accounts.ValidateAccount(ctx, req.ToAccountID, req.Currency)
  if err != nil {
		m.log.Errorf("m.accounts.ValidateAccount: %v", err)
		return nil, errors.Wrapf(err, "m.TransferTransaction")
  }

  if !*isToAccountValid {
    return nil, errors.Wrapf(err, "m.TransferTransaction")
  }

  // Initiate transfer
	result, err := m.db.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		m.log.Errorf("m.db.CreateTransfer: %v", err)
		return nil, errors.Wrapf(err, "m.TransferTransaction")
	}

	// TransferTxResult is a pretty complicated struct,
	// Will not do a module layer for the type for now
	return &result, nil
}

func (m *Module) ValidateAccount(ctx context.Context, accountID int64, currency string) (*bool, error) {
  return m.accounts.ValidateAccount(ctx, accountID, currency)
}
