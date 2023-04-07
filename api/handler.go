package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

func (s *Server) createAccount(ctx *gin.Context) {
	var req CreateAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Info(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(BadRequest))
		return
	}

	acc, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		OwnerID:  req.OwnerID,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, toResponseBody(convertToAccountResponse(acc)))
}

func (s *Server) getAccount(ctx *gin.Context) {
	pathID := ctx.Param("id")
	if pathID == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(BadRequest))
		return
	}

	id, err := strconv.ParseInt(pathID, 10, 64)
  if err != nil {
    s.log.Info(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(BadRequest))
		return
  }

  acc, err := s.store.GetAccount(ctx, id)
  if err != nil {
    s.log.Info("s.store.GetAccount", err.Error())
    ctx.JSON(http.StatusInternalServerError, nil)
  }

  ctx.JSON(http.StatusOK, toResponseBody(convertToAccountResponse(acc)))
}
