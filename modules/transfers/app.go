package transfers

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	accountsAdapter "github.com/muazwzxv/go-backend-masterclass/modules/transfers/adapters/accounts"
	adapter "github.com/muazwzxv/go-backend-masterclass/modules/transfers/adapters/accounts"
	"go.uber.org/zap"
)

type ITransfers interface {
	TransferTransaction(ctx context.Context, req *TransferRequest) (*db.TransferTxResult, error)
}

type Module struct {
	db       db.IStore
	log      *zap.SugaredLogger
	accountsAdapter accountsAdapter.IAccounts
}

func New(
	db db.IStore,
	log *zap.SugaredLogger,
	accountsAdapter accountsAdapter.IAccounts,
) *Module {
	return &Module{
		db:       db,
		log:      log,
		accountsAdapter: accountsAdapter,
	}
}

// verify transfers.Module implements ITransfers
var (
	_ ITransfers = (*Module)(nil)
  _ adapter.IAccounts = (*adapter.AccountsAdapter)(nil)
  _ adapter.IAccounts = (*accounts.Module)(nil)
)
