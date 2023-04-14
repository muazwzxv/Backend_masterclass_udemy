package db

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type IStore interface {
  Querier
	TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error)
}

var _ IStore = (*Store)(nil)

// provides function to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// exec a function within a database transactions
func (s *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "s.execTX")
	}
  queries := New(tx)


	err = fn(queries)
	if err != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return errors.Wrapf(err, "rollback Error %v", rollbackError)
		}
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performs money transfer from one account to another
// Creates transfer record
// Add account entries
// Update accounts balance
func (s *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var res TransferTxResult
	err := s.execTX(ctx, func(q *Queries) error {
		var err error

		res.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "s.transferTx.CreateTransfer")
		}

		res.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: sql.NullInt64{
				Int64: args.FromAccountID,
				Valid: true,
			},
			Amount: args.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "s.transferTx.CreateEntry")
		}

		res.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: sql.NullInt64{
				Int64: args.ToAccountID,
				Valid: true,
			},
			Amount: args.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "s.transferTx.CreateEntry")
		}

		res.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.FromAccountID,
			Amount: args.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "s.transferTx.UpdateAccountBalance")
		}

		res.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.ToAccountID,
			Amount: args.Amount,
		})
		if err != nil {
			return errors.Wrap(err, "s.transferTx.UpdateAccountBalance")
		}

		return nil
	})

	return res, err
}
