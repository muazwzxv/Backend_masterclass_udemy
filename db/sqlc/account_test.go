package db_test

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	userArg := db.CreateUserParams{
		FirstName: sql.NullString{String: "name", Valid: true},
		LastName:  sql.NullString{String: "default", Valid: true},
		Email:     "default@gmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), userArg)
	require.NoError(t, err)

	accountArg := db.CreateAccountParams{
		OwnerID:  user.ID,
		Balance:  1000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), accountArg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, accountArg.OwnerID, account.OwnerID)
	require.Equal(t, accountArg.Balance, account.Balance)
	require.Equal(t, accountArg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
