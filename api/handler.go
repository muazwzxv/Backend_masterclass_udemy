package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

func (s *Server) createAccount(ctx *gin.Context) {
	var req CreateAccount
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  acc, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
    OwnerID: req.OwnerID,
    Balance: 0,
    Currency: req.Currency,
  })

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, errorResponse(err)) 
    return
  }
  ctx.JSON(http.StatusCreated, toResponseBody(convertToAccountResponse(acc)))
}
