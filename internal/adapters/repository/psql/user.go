package psql

import (
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

func (u *userRepository) CreateUser() {}
func (u *userRepository) LoginUser()  {}
