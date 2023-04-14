package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
)

type Gateway struct {
	accounts *accounts.Handler
  users *users.Handler
}

func New(
	accounts *accounts.Handler,
  users *users.Handler,
) *Gateway {
	return &Gateway{
		accounts: accounts,
    users: users,
	}
}

func (g *Gateway) Init(mux *gin.Engine) {
	api := mux.Group("/api")

	// setup routes
	g.accounts.Routes(api)
  g.users.Routes(api)
}
