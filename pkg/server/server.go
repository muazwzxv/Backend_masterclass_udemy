package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type Server struct {
	Store db.IStore
	Mux   *gin.Engine
	Log   *zap.SugaredLogger
}

var _ IServer = (*Server)(nil)

func NewServer(
	store db.IStore,
	log *zap.SugaredLogger,
) *Server {
	mux := gin.Default()

	server := &Server{
		Store: store,
		Mux:   mux,
		Log:   log,
	}
	return server
}

func (s *Server) Start(address string) error {
	return s.Mux.Run(address)
}

func (s *Server) Stop() error {
	return nil
}
