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
	GetAccount(username string) (*models.Account, error)
	CreateAccount(username, email, password string) (int64, error)
	CreateActivationToken(accountID int64, activationToken string) (int64, error)
	ActivateAccount(token string) (bool, error)
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
func (a *authRepository) GetAccount(username string) (*models.Account, error) {
	account := models.Account{}

	err := a.DB.Get(&account, `SELECT
		*
		FROM accounts a
		WHERE a.username = ?;
	`,
		username,
	)
	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil
}

// CreateAccount creates an account with the given username, email, and password.
func (a *authRepository) CreateAccount(username, email, password string) (int64, error) {
	result, err := a.DB.Exec(`INSERT
		INTO accounts(
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

// CreateActivationToken creates an activation token for the given account ID.
func (a *authRepository) CreateActivationToken(accountID int64, activationToken string) (int64, error) {
	result, err := a.DB.Exec(`INSERT
		INTO account_activation_tokens(
			account_id,
			activation_token
		) VALUES (
			?,
			?
		);
	`,
		accountID,
		activationToken,
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

// ActivateAccount marks the account with the given token as active and expires the token.
func (a *authRepository) ActivateAccount(token string) (bool, error) {
	var accountID int64
	err := a.DB.Get(&accountID, `SELECT
		a.id
		FROM accounts a
		INNER JOIN account_activation_tokens t ON a.id = t.account_id
		WHERE t.activation_token = ?
		AND t.is_used = 0;
	`,
		token,
	)
	if err != nil {
		return false, err
	}

	tx, err := a.DB.Begin()
	_, err = tx.Exec(`UPDATE account_activation_tokens
		SET is_used = 1
		WHERE account_id = ?
		AND activation_token = ?;
	`,
		accountID,
		token,
	)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec(`UPDATE accounts
		SET is_activated = 1
		WHERE id = ?;
	`,
		accountID,
	)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}
