package transfers

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/modules/transfers"
	"go.uber.org/zap"
)

type Handler struct {
  m transfers.ITransfers
	log *zap.SugaredLogger
}

func New(
  module transfers.ITransfers,
  log *zap.SugaredLogger,
) *Handler {
  return &Handler{
    m: module,
    log: log,
  }
}

func (h *Handler) Routes(route *gin.RouterGroup) {
	v1 := route.Group("/v1")
	{
		v1.GET("/transfer", func(ctx *gin.Context) {
		})
	}
}
