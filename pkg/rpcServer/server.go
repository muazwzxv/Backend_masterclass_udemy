package rpcServer

import (
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config *config.Config
	store  db.IStore
	token  authToken.IToken
}

func NewServer(cfg *config.Config, store *db.Store, token authToken.IToken) *Server {
	server := &Server{
		config: cfg,
		store:  store,
		token:  token,
	}
	return server
}
