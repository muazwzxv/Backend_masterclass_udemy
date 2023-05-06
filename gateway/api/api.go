package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/transfers"
	"github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
)

type Gateway struct {
	accounts  *accounts.Handler
	users     *users.Handler
	transfers *transfers.Handler
}

func New(
	accounts *accounts.Handler,
	users *users.Handler,
	transfers *transfers.Handler,
) *Gateway {
	return &Gateway{
		accounts:  accounts,
		users:     users,
		transfers: transfers,
	}
}

func (g *Gateway) Init(mux *gin.Engine) {
	api := mux.Group("/api")

  // register custom validator
	g.registerCustomValidator(mux)

	// setup routes
	g.accounts.Routes(api)
	g.users.Routes(api)
	g.transfers.Routes(api)
}

func (g *Gateway) registerCustomValidator(mux *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    err := v.RegisterValidation("currency", transfers.ValidCurrency)
    if err != nil {
      log.Fatalf("failed to register vaidation")
    }
	}
}
