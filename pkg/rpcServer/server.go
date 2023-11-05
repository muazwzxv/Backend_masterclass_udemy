package rpcServer

import (
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Store  db.IStore
	Log    *zap.SugaredLogger
	Token  authToken.IToken
}

// TODO - implement the server adapter
func NewServer(cfg *config.Config, store *db.Store, log *zap.SugaredLogger, token authToken.IToken) *Server {
	server := &Server{
		Config: cfg,
		Store:  store,
		Log:    log,
		Token:  token,
	}
	return server
}
