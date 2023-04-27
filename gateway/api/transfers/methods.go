package transfers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/utils"
)

func (h *Handler) CreateTransfer(ctx *gin.Context) {
  var req TransferRequest
  if err := ctx.ShouldBindJSON(&req); err != nil {
    h.log.Errorf("h.CreateTransfer", err)
    ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
    return
  }

  // TODO: Module layer for TransferTx
}
