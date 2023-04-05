package db_test

import (
	"context"
	"testing"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := db.NewStore(testDB)

	account1 := db.Account{
		ID:       1,
		OwnerID:  1,
		Balance:  200,
		Currency: "USD",
	}
	account2 := db.Account{
		ID:       2,
		OwnerID:  1,
		Balance:  400,
		Currency: "USD",
	}

	// run a concurrent transfer transaction
	n := 5
	amount := int64(20)

	errs := make(chan error)
	result := make(chan db.TransferTxResult)

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		go func() {
			res, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: int64(account1.ID),
				ToAccountID:   int64(account2.ID),
				Amount:        amount,
			})

			// send error to the main goroutine
			errs <- err
			result <- res
		}()
	}

	// Assert result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		res := <-result
		require.NotEmpty(t, res)

		// check transfers
		transfer := res.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entry from
		fromEntry := res.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// check entry to
		toEntry := res.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account1.ID, toEntry.AccountID)
		require.Equal(t, -amount, toEntry.Amount)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := res.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := res.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check account's balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check final updated balance
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
