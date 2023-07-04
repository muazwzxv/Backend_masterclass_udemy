package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	type testcase struct {
		Title       string
		Password    string
		ExpectedErr error
	}

	testCases := []testcase{
		{
			Title:       "Password is properly given",
			Password:    "The password lmaoo",
			ExpectedErr: nil,
		},
		{
			Title:       "Password is empty string",
			Password:    "",
			ExpectedErr: ErrPasswordIsEmpty,
		},
	}

	for _, tc := range testCases {
		_, err := HashPassword(tc.Password)
		require.Equal(t, tc.ExpectedErr, err)
	}
}

func TestCheckPassword(t *testing.T) {
	type testcase struct {
		Title       string
		Hashed      string
		Password    string
		ExpectedErr error
	}

	testCases := []testcase{
		{
			Title:       "Check password successful",
			Hashed:      "$2a$10$A.jSUk9RYIt.eXgSPP4JGuD1WwaU1UuQRyOE9FW4aEfnmhleZEI2e,",
			Password:    "The password lmaoo",
			ExpectedErr: nil,
		},
		{
			Title:       "Password does not match",
			Hashed:      "$2a$10$A.jSUk9RYIt.eXgSPP4JGuD1WwaU1UuQRyOE9FW4aEfnmhleZEI2e,",
			Password:    "Another one ",
			ExpectedErr: bcrypt.ErrMismatchedHashAndPassword,
		},
	}

	for _, tc := range testCases {
		err := CheckPassword(tc.Password, tc.Hashed)
		require.ErrorIs(t, tc.ExpectedErr, err)
	}
}
