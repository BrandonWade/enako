package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_auth_repository.go . AuthRepository
type AuthRepository interface {
	CreateAccount(username, email, password string) (int64, error)
	GetAccount(username string) (*models.UserAccount, error)
}

type authRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DB *sqlx.DB) AuthRepository {
	return &authRepository{
		DB,
	}
}

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
