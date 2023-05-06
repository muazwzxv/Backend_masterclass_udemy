package transfers

const (
	USD = "USD"
	EUR = "EUR"
	MYR = "MYR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MYR:
		return true
	}
  return false
}

type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}
