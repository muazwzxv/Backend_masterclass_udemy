package users

// TODO: create API layer gateway
import usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"

func convertToResponseUser(user *usersModule.User) *User {
  return &User{
    ID: user.ID,
    FirstName: user.FirstName,
    LastName: user.LastName,
    Email: user.Email,
    CreatedAt: user.CreatedAt,
    DeletedAt: user.DeletedAt,
  }
}
