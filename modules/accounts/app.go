package accounts

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type IAccounts interface {
	CreateAccount(ctx context.Context, data *CreateAccount) (*db.Account, error)
	FindAccount(ctx context.Context, id int64) (*db.Account, error)
}

type Module struct {
	db  *db.Store
	log *zap.SugaredLogger // should this be here or handler layer? (i feel like handler and not here)
}

func New(
  db *db.Store, 
  log *zap.SugaredLogger,
) *Module {
	return &Module{
		db: db,
	}
}

// verify implementor
var _ IAccounts = (*Module)(nil)