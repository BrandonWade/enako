package repositories

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_auth_repository.go . AuthRepository
type AuthRepository interface {
	CreateAccount(email, password string) (int64, error)
}

type authRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DB *sqlx.DB) AuthRepository {
	return &authRepository{
		DB,
	}
}

func (a *authRepository) CreateAccount(email, password string) (int64, error) {
	result, err := a.DB.Exec(`INSERT
		INTO user_accounts(
			email,
			password
		) VALUES (
			?,
			?
		);
	`,
		email,
		password,
	)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID, nil
}
