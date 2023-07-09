package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/utils"
	dbErr "github.com/muazwzxv/go-backend-masterclass/db/errors"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req CreateAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("h.CreateAccount: %v", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	acc, err := h.m.CreateAccount(ctx, &accountsModule.CreateAccount{
		OwnerID:  req.OwnerID,
		Currency: req.Currency,
		Balance:  0,
	})
	if err != nil {
		errCode := dbErr.ErrorCode(err)
		if errCode == dbErr.ForeignKeyViolation || errCode == dbErr.UniqueViolation {
      ctx.AbortWithStatus(http.StatusForbidden)
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, utils.ToResponseBody(acc))
}

func (h *Handler) GetAccount(ctx *gin.Context) {
	var req GetAccountrequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	acc, err := h.m.FindAccount(ctx, req.ID)
	if err != nil {
		h.log.Errorf("h.GetAccount: %v", err)
		if err == accountsModule.ErrNotFound {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ToResponseBody(acc))
}

func (h *Handler) ListAccounts(ctx *gin.Context) {
	var req GetAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.log.Errorf("h.ListAccounts", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	accs, err := h.m.ListAccounts(ctx, &accountsModule.GetAccounts{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, utils.ToResponseBody(accs))
}
