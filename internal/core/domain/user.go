package domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	BaseModel
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	HashedPassword []byte `json:"-,omitempty"`
}

type UserResponse struct {
	ID        int64
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
