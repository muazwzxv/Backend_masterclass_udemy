package users

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"


func convertToModuleUser(usr *db.User) *User {
  return &User{
    ID: usr.ID,
    FirstName: usr.FirstName.String,
    LastName: usr.LastName.String,
    Email: usr.Email,
    CreatedAt: &usr.CreatedAt.Time,
    DeletedAt: &usr.DeletedAt.Time,
  }
}
