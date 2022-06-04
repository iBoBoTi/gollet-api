package domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrExpiredToken    = errors.New("expired token")
	ErrInvalidPassword = bcrypt.ErrMismatchedHashAndPassword
)
