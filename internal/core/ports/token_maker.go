package ports

import (
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"time"
)

type TokenMaker interface {
	CreateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*domain.Payload, error)
}
