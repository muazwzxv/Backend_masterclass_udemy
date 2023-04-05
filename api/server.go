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
  server := &Server{
    store: store,
    mux: gin.Default(),
  }

  return server
}
