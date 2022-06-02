package psql

import (
	"context"
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
	queryString := `INSERT INTO users (email, name, hashed_password) VALUES ($1, $2, $3) RETURNING *`
	result := &domain.User{}
	row := u.db.QueryRow(context.Background(), queryString, user.Email, user.Name, user.HashedPassword)
	err := row.Scan(&result.ID, &result.Name, &result.Email, &result.HashedPassword, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) CreateUserWallet(wallet *domain.Wallet) (*domain.Wallet, error) {
	queryString := `INSERT INTO wallets (user_id, balance, currency) VALUES ($1, $2, $3) RETURNING *`
	result := &domain.Wallet{}
	row := u.db.QueryRow(context.Background(), queryString, wallet.UserID, wallet.Balance, wallet.Currency)
	err := row.Scan(&result.ID, &result.UserID, &result.Balance, &result.Currency, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) LoginUser() {}
