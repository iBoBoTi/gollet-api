package usecase

import (
	"errors"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"time"
)

type userService struct {
	userRepository ports.UserRepository
	tokenMaker     ports.TokenMaker
}

func NewUserService(userRepo ports.UserRepository, tokenMaker ports.TokenMaker) ports.UserService {
	return &userService{
		userRepository: userRepo,
		tokenMaker:     tokenMaker,
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

func (u *userService) LoginUser(loginRequest *domain.LoginUserRequest) (*domain.LoginUserResponse, error) {
	foundUser, err := u.userRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	err = foundUser.VerifyPassword(loginRequest.Password)
	if err != nil {
		return nil, err
	}

	userResponse := domain.UserResponse{
		ID:        foundUser.ID,
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		CreatedAt: foundUser.CreatedAt,
		UpdatedAt: foundUser.UpdatedAt,
	}
	token, err := u.tokenMaker.CreateToken(foundUser.Email, 3*time.Duration(24)*time.Minute)
	if err != nil {
		return nil, err
	}

	return &domain.LoginUserResponse{
		UserResponse: userResponse,
		AccessToken:  token,
	}, nil
}
