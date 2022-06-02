package domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required, email"`
	Password       string `json:"password" binding:"required"`
	HashedPassword []byte `json:"-,omitempty"`
}

func (u *User) SetPassword(password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.HashedPassword = hash
	return nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(password string) error {
	if u.HashedPassword != nil && len(u.HashedPassword) == 0 {
		// Internal Server
		return errors.New("password is not set")
	}
	// Wrong Password
	return bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password))
}
