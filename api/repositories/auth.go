package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// AuthRepository an interface for working with accounts.
//go:generate counterfeiter -o fakes/fake_auth_repository.go . AuthRepository
type AuthRepository interface {
	GetAccount(username string) (*models.UserAccount, error)
	CreateAccount(username, email, password string) (int64, error)
}

type authRepository struct {
	DB *sqlx.DB
}

// NewAuthRepository returns a new instance of an AuthRepository.
func NewAuthRepository(DB *sqlx.DB) AuthRepository {
	return &authRepository{
		DB,
	}
}

// GetAccount returns an account with the given username.
func (a *authRepository) GetAccount(username string) (*models.UserAccount, error) {
	account := models.UserAccount{}

	err := a.DB.Get(&account, `SELECT
		*
		FROM user_accounts u
		WHERE u.username = ?
	`,
		username,
	)
	if err != nil {
		return &models.UserAccount{}, err
	}

	return &account, nil
}

// CreateAccount creates an account with the given username, email, and password.
func (a *authRepository) CreateAccount(username, email, password string) (int64, error) {
	result, err := a.DB.Exec(`INSERT
		INTO user_accounts(
			username,
			email,
			password
		) VALUES (
			?,
			?,
			?
		);
	`,
		username,
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
