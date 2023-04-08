package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
)

type Handler struct {
	m   *accounts.Module
}

func New(
	module *accounts.Module,
) *Handler {
	return &Handler{
		m:   module,
	}
}

func (h *Handler) Routes(route *gin.RouterGroup) {
	v1 := route.Group("/v1")
	{
		v1.POST("/accounts", h.CreateAccount)
		v1.GET("/accounts/:id", h.GetAccount)
	}
}
