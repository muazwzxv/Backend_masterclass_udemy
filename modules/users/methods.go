package users

import (
	"context"
	"database/sql"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/pkg/hash"
	"github.com/pkg/errors"
)

func (m *Module) CreateUser(ctx context.Context, data *CreateUser) (*User, error) {
	hashed, err := hash.HashPassword(data.Password)
	if err != nil {
		m.log.Error("m.db.CreateUser: %v", err)
		return nil, errors.Wrap(err, "hash.HashPassword")
	}

	user, err := m.db.CreateUser(ctx, db.CreateUserParams{
		FirstName:      sql.NullString{String: data.FirstName, Valid: true},
		LastName:       sql.NullString{String: data.LastName, Valid: true},
		UserName:       data.UserName,
		Email:          data.Email,
		HashedPassword: hashed,
	})
	if err != nil {
		m.log.Errorf("m.db.CreateUser: %v", err)
		return nil, errors.Wrapf(err, "m.CreateUser")
	}

	return convertToModuleUser(&user), nil
}

func (m *Module) FindUser(ctx context.Context, id int64) (*User, error) {
	user, err := m.db.GetUser(ctx, id)
	if err != nil {
		m.log.Errorf("m.db.Getusers: %v", err)
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrapf(err, "m.FindUser")
	}
	return convertToModuleUser(&user), nil
}

func (m *Module) LoginUser(ctx context.Context, data *LoginUserRequest) (*LoginResponse, error) {
	user, err := m.db.GetUserByUsername(ctx, data.UserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, ErrNotFound.Error())
		}
		return nil, errors.Wrap(err, ErrFailedToQueryUser.Error())
	}

	// TODO - check password

	// Create token
	userToken, err := m.token.CreateToken(data.UserName, m.Config.AccessTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, ErrFaileToGenerateToken.Error())
	}

	res := &LoginResponse{
		Token:    userToken,
		UserData: *convertToModuleUser(&user),
	}

	return res, nil
}

func (m *Module) UpdateUser(ctx context.Context, data *UpdateUser) error {
	// TODO:
	return nil
}

func (m *Module) UpdatePassword(ctx context.Context, data *UpdatePassword) error {
	// TODO:
	return nil
}
