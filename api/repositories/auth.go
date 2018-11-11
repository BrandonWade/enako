package repositories

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type AuthRepository interface {
	CreateAccount(email, password string) error
}

type authRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DB *sqlx.DB) AuthRepository {
	return &authRepository{
		DB,
	}
}

func (a *authRepository) CreateAccount(email, password string) error {
	return nil
}
