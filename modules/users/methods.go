package users

import (
	"context"
	"database/sql"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/pkg/errors"
)

func (m *Module) CreateUser(ctx context.Context, data *CreateUser) (*User, error) {
	user, err := m.db.CreateUser(ctx, db.CreateUserParams{
    FirstName: sql.NullString{String: data.FirstName, Valid: true},
    LastName: sql.NullString{String: data.LastName, Valid: true},
    Email: data.Email,
  })
  if err != nil {
    m.log.Errorf("m.db.CreateUser: %v", err)
    return nil, errors.Wrapf(err, "m.CreateUser")
  }

	return convertToModuleUser(&user), nil
}

func (m *Module) FindUser(ctx context.Context, id int64) (*User, error) {
  user, err := m.db.GetUsers(ctx, id)
  if err != nil {
    m.log.Errorf("m.db.Getusers: %v", err)
    if err == sql.ErrNoRows {
      return nil, NotFound
    }
    return nil, errors.Wrapf(err, "m.FindUser")
  }
  return convertToModuleUser(&user), nil
}

func (m *Module) UpdateUser(ctx context.Context, data *UpdateUser) error {
  // TODO:
  return nil
}


func (m *Module) UpdatePassword(ctx context.Context, data *UpdatePassword) error {
  // TODO:
  return nil
}

