package usecase

import "github.com/iBoBoTi/gollet-api/internal/core/ports"

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (u *userService) CreateUser() {}

func (u *userService) LoginUser() {}
