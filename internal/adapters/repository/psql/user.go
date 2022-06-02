package psql

import (
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) SignUpUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) LoginUser() {}
