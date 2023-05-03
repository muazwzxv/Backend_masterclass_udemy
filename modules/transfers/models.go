package transfers

type TransferRequest struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	Currency      string
}
