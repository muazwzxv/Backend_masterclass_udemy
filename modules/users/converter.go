package users

import db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"

func convertToModuleUser(usr *db.User) *User {
	user := &User{
		ID:        usr.ID,
		FirstName: usr.FirstName.String,
		LastName:  usr.LastName.String,
		UserName:  usr.UserName,
		Email:     usr.Email,
		Password:  usr.HashedPassword,
		CreatedAt: &usr.CreatedAt.Time,
		DeletedAt: &usr.DeletedAt.Time,
	}

	if !usr.CreatedAt.Valid {
		user.CreatedAt = nil
	}

	if !usr.DeletedAt.Valid {
		user.DeletedAt = nil
	}

	return user
}
