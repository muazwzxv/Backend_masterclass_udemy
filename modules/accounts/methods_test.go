package accounts_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/muazwzxv/go-backend-masterclass/db/mock"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetAccountAPI(t *testing.T) {
	testAcc := randomAccount()

	// create mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockIStore(ctrl)

	// mock
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
		Times(1).
		Return(testAcc, nil)

	type testCase struct {
		Name        string
		Account     db.Account
		ExpectedErr error
	}

	testCases := []testCase{
		{
			Name:        "Account found",
			Account:     randomAccount(),
			ExpectedErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			module := createModule(store)
			acc, err := module.FindAccount(context.Background(), testAcc.ID)
      require.NoError(t, err)
      require.NotNil(t, acc)
		})
	}
}

// Helpers

func createModule(db db.IStore) accounts.IAccounts {
	log, _ := zap.NewDevelopment()
	sugaredLogger := log.Sugar()

	return accounts.New(db, sugaredLogger)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       rand.Int63(),
		OwnerID:  rand.Int63(),
		Balance:  200,
		Currency: db.CurrencyEUR,
	}
}
