package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
)

type Gateway struct {
  accounts *accounts.Handler 
}

func New(
  accounts *accounts.Handler,
) *Gateway {
  return &Gateway{
    accounts: accounts,
  }
}
 
func (g *Gateway) Init(mux *gin.Engine) {
  api := mux.Group("/api")

  // setup routes
  g.accounts.Routes(api)
}


