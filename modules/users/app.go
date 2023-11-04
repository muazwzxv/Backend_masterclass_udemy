package users

import (
	"context"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"go.uber.org/zap"
)

type IUsers interface {
	CreateUser(ctx context.Context, data *CreateUser) (*User, error)
	FindUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, data *UpdateUser) error
	UpdatePassword(ctx context.Context, data *UpdatePassword) error
	LoginUser(ctx context.Context, data *LoginUserRequest) (*LoginResponse, error)
}

type Module struct {
	db     db.IStore
	log    *zap.SugaredLogger
	token  authToken.IToken
	Config *config.Config
}

func New(
	config *config.Config,
	db db.IStore,
	log *zap.SugaredLogger,
	token authToken.IToken,
) *Module {
	return &Module{
		db:     db,
		log:    log,
		token:  token,
		Config: config,
	}
}

// verify users.Module implements IAccounts
var _ IUsers = (*Module)(nil)
