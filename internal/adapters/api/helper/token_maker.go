package helper

import (
	"errors"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
)

const minSecretKeyLength = 32

func NewJWTMaker(secretKey string) (ports.TokenMaker, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, errors.New("secret key is too short")
	}
	return &domain.JWTMaker{SecretKey: secretKey}, nil
}
