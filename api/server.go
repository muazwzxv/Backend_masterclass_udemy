package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

type Server struct {
	store *db.Store
	mux   *gin.Engine
}

func NewServer(store *db.Store) *Server {
  mux := gin.Default()
  server := &Server{
    store: store,
    mux: mux,
  }
  return server
}

func (s *Server) Start(address string) error {
  return s.mux.Run(address)
}

func (s *Server) SetupRoutes() {
  s.mux.POST("/accounts", s.createAccount)
}
