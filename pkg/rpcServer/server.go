package rpcServer

import (
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"github.com/muazwzxv/go-backend-masterclass/pkg/worker"
	"go.uber.org/zap"
)

type Server struct {
	Config          *config.Config
	Store           db.IStore
	Log             *zap.SugaredLogger
	Token           authToken.IToken
	TaskDistributor worker.TaskDistributor
}

type ServerRequest struct {
	Cfg             *config.Config
	Store           *db.Store
	Log             *zap.SugaredLogger
	Token           authToken.IToken
	TaskDistributor worker.TaskDistributor
}

func NewServer(req ServerRequest) *Server {
	server := &Server{
		Config:          req.Cfg,
		Store:           req.Store,
		Log:             req.Log,
		Token:           req.Token,
		TaskDistributor: req.TaskDistributor,
	}
	return server
}
