package transfers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/utils"
	"github.com/muazwzxv/go-backend-masterclass/modules/transfers"
)

func (h *Handler) CreateTransfer(ctx *gin.Context) {
	var req TransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("h.CreateTransfer", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

  res, err := h.m.TransferTransaction(ctx, &transfers.TransferRequest{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	})
  if err != nil {
    h.log.Info("i made it in internal server error")
    ctx.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  ctx.JSON(http.StatusCreated, utils.ToResponseBody(res))
}
