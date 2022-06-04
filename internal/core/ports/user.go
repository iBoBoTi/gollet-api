package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
)

type UserHandler interface {
	SignUpUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserService interface {
	SignUpUser(user *domain.User) (*domain.User, error)
	LoginUser(user *domain.LoginUserRequest) (*domain.LoginUserResponse, error)
}

type UserRepository interface {
	SignUpUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	CreateUserWallet(wallet *domain.Wallet) (*domain.Wallet, error)
}
