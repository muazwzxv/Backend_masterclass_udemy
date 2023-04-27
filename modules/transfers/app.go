package transfers

import (
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type ITransfers interface{
}

type Module struct {
	db  db.IStore
	log *zap.SugaredLogger
}

func New(
	db db.IStore,
	log *zap.SugaredLogger,
) ITransfers {
	return &Module{
		db:  db,
		log: log,
	}
}

// verify transfers.Module implements ITransfers
var _ ITransfers = (*Module)(nil)
