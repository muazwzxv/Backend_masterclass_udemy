package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"go.uber.org/zap"
)

type Server struct {
	store *db.Store
	mux   *gin.Engine
  log  *zap.SugaredLogger
}

func NewServer(
  store *db.Store, 
  log *zap.SugaredLogger,
) *Server {
	mux := gin.Default()

	server := &Server{
		store: store,
		mux:   mux,
    log: log,
	}
	return server
}

func (s *Server) Start(address string) error {
	return s.mux.Run(address)
}

func (s *Server) SetupRoutes() {
	api := s.mux.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/accounts", s.createAccount)
		}
	}
}
