package usecase

import (
	"errors"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (u *userService) SignUpUser(user *domain.User) (*domain.User, error) {
	// ToDo: create user
	// Todo: create user wallet
	// Todo: if wallet creation fails, rollback user creation
	// signup user repository in the case of no error call create wallet repository if no error return user else return error
	// hash user password
	err := user.SetPassword(user.Password)
	if err != nil {
		return nil, err
	}

	result, err := u.userRepository.SignUpUser(user)
	if err != nil {
		return nil, err
	}
	// Then create user wallet
	var wallet = domain.Wallet{
		UserID:   result.ID,
		Balance:  0,
		Currency: "NGN",
	}
	_, err = u.userRepository.CreateUserWallet(&wallet)
	if err != nil {
		err = errors.New("wallet creation error")
		return nil, err
	}
	return result, nil
}

func (u *userService) LoginUser() {}
