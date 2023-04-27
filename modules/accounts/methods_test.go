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

	testCases := []struct {
		Name        string
		AccountID   int64
		ExpectedErr error
		buildStubs  func(store *mockdb.MockIStore)
	}{
		{
			Name:        "Account found",
			AccountID:   testAcc.ID,
			ExpectedErr: nil,
			buildStubs: func(store *mockdb.MockIStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
					Times(1).
					Return(testAcc, nil)
			},
		},
		{
			Name:        "Account Not found",
			AccountID:   20,
			ExpectedErr: accounts.NotFound,
			buildStubs: func(store *mockdb.MockIStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
		},
		{
			Name:        "Internal server error",
			AccountID:   20,
			ExpectedErr: sql.ErrConnDone,
			buildStubs: func(store *mockdb.MockIStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.buildStubs(store)
			_, err := module.FindAccount(context.Background(), tc.AccountID)
			assert.ErrorIs(t, err, tc.ExpectedErr)
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
