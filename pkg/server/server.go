package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	Store  db.IStore
	Mux    *gin.Engine
	Log    *zap.SugaredLogger
	Token  authToken.IToken
	Config *config.Config
}

var _ IServer = (*Server)(nil)

type HttpServerRequest struct {
	Config *config.Config
	Store  db.IStore
	Log    *zap.SugaredLogger
	Token  authToken.IToken
}

func NewServer(req HttpServerRequest) *Server {
	mux := gin.Default()

	server := &Server{
		Store: req.Store,
		Mux:   mux,
		Log:   req.Log,
		Token: req.Token,
	}
	return server
}

func (s *Server) Start(address string) error {
	return s.Mux.Run(address)
}

func (s *Server) Stop() error {
	return nil
}
