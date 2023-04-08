package accounts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/utils"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req CreateAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	acc, err := h.m.CreateAccount(ctx, &accountsModule.CreateAccount{
		OwnerID:  req.OwnerID,
		Currency: req.Currency,
		Balance:  0,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.ToResponseBody(convertToAccountResponse(acc)))
}

func (h *Handler) GetAccount(ctx *gin.Context) {
	pathID := ctx.Param("id")
	if pathID == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	id, err := strconv.ParseInt(pathID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

  acc, err := h.m.FindAccount(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, utils.ToResponseBody(convertToAccountResponse(acc)))
}
