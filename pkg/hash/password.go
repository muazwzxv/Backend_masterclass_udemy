package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
  hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", fmt.Errorf("failed to hash password %w", err)
  }
  return string(hashed), nil
}

func CheckPassword(password string, hashedPass string) error {
  return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPass))
}
