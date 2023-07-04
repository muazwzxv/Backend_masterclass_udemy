package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
  if password == "" {
    return "", ErrPasswordIsEmpty
  }

  hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", fmt.Errorf("failed to hash password %w", err)
  }
  return string(hashed), nil

}
func CheckPassword(password string, hashedPass string) error {
  if password == "" || hashedPass == "" {
    return ErrPasswordIsEmpty
  }

  return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
