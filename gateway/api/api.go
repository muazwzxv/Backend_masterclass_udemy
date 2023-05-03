package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/transfers"
)

type Gateway struct {
	accounts *accounts.Handler
  users *users.Handler
  transfers *transfers.Handler
}

func New(
	accounts *accounts.Handler,
  users *users.Handler,
  transfers *transfers.Handler,
) *Gateway {
	return &Gateway{
		accounts: accounts,
    users: users,
    transfers: transfers,
	}
}

func (g *Gateway) Init(mux *gin.Engine) {
	api := mux.Group("/api")

	// setup routes
	g.accounts.Routes(api)
  g.users.Routes(api)
  g.transfers.Routes(api)
}
