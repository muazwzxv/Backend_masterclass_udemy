package accounts_test

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/muazwzxv/go-backend-masterclass/db/mock"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetAccount(t *testing.T) {
	testAcc := randomAccount()

	// create mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockIStore(ctrl)

	// create module
	module := createModule(store)

  // setup mock
  store.EXPECT().
    GetAccount(gomock.Any(), gomock.Any()).
    DoAndReturn(func(ctx context.Context, ID int64) (db.Account, error) {
      if ID == testAcc.ID {
        return testAcc, nil
      } 
      return db.Account{}, sql.ErrNoRows
    }).
  AnyTimes()

	type testCase struct {
		Name        string
		AccountID   int64
		ExpectedErr error
	}

	testCases := []testCase{
		{
			Name:        "Account found",
			AccountID:   testAcc.ID,
			ExpectedErr: nil,
		},
		{
			Name:        "Account Not found",
			AccountID:   20,
			ExpectedErr: accounts.NotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			_, err := module.FindAccount(context.Background(), test.AccountID)
			assert.ErrorIs(t, test.ExpectedErr, err)
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
