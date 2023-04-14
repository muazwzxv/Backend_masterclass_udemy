package pkg

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type Server struct {
	Store *db.Store
	Mux   *gin.Engine
	Log   *zap.SugaredLogger
}

func NewServer(
	store *db.Store,
	log *zap.SugaredLogger,
) *Server {
	mux := gin.New()

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
