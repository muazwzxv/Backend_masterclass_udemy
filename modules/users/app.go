package users

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type IUsers interface{
  CreateUser(ctx context.Context, data *CreateUser) (*User, error)
  FindUser(ctx context.Context, id int64) (*User, error)
}

type Module struct {
	db  db.IStore
	log *zap.SugaredLogger
}

func New(
  db db.IStore,
  log *zap.SugaredLogger,
) *Module {
  return &Module{
    db: db,
    log: log,
  }
}

// verify users.Module implements IAccounts
var _ IUsers = (*Module)(nil)
