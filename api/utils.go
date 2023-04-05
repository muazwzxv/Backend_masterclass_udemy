package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func toResponseBody(data interface{}) gin.H {
	return gin.H{"data": data}
}

func convertToAccountResponse(acc db.Account) Account {
	return Account{
		ID:       acc.ID,
		OwnerID:  acc.OwnerID,
		Balance:  acc.Balance,
		Currency: acc.Currency,
	}
}
