package transfers

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	"go.uber.org/zap"
)

type ITransfers interface {
	IAccounts
	TransferTransaction(ctx context.Context, req *TransferRequest) (*db.TransferTxResult, error)
}

type IAccounts interface {
	ValidateAccount(ctx context.Context, accountID int64, currency string) (*bool, error)
}

type Module struct {
	db       db.IStore
	log      *zap.SugaredLogger
	accounts IAccounts
}

func New(
	db db.IStore,
	log *zap.SugaredLogger,
	accountsModule IAccounts,
) *Module {
	return &Module{
		db:       db,
		log:      log,
		accounts: accountsModule,
	}
}

// verify transfers.Module implements ITransfers
var (
	_ ITransfers = (*Module)(nil)
	_ IAccounts  = (*accounts.Module)(nil)
)
