package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// AccountRepository an interface for working with accounts.
//go:generate counterfeiter -o fakes/fake_account_repository.go . AccountRepository
type AccountRepository interface {
	GetAccount(username string) (*models.Account, error)
	CreateAccount(username, email, password string) (int64, error)
	CreateActivationToken(accountID int64, activationToken string) (int64, error)
	ActivateAccount(token string) (bool, error)
	GetAccountByUsername(username string) (*models.Account, error)
	CreatePasswordResetToken(accountID int64, resetToken string) (int64, error)
	GetPasswordResetToken(token string) (*models.PasswordResetToken, error)
	ResetPassword(token, password string) (bool, error)
}

type accountRepository struct {
	DB *sqlx.DB
}

// NewAccountRepository returns a new instance of an AccountRepository.
func NewAccountRepository(DB *sqlx.DB) AccountRepository {
	return &accountRepository{
		DB,
	}
}

// GetAccount returns an account with the given username.
func (a *accountRepository) GetAccount(username string) (*models.Account, error) {
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
func (a *accountRepository) CreateAccount(username, email, password string) (int64, error) {
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
func (a *accountRepository) CreateActivationToken(accountID int64, activationToken string) (int64, error) {
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
func (a *accountRepository) ActivateAccount(token string) (bool, error) {
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

// GetAccountByUsername returns the account with the given username.
func (a *accountRepository) GetAccountByUsername(username string) (*models.Account, error) {
	var account models.Account

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

// CreatePasswordResetToken creates an activation token for the given account ID.
func (a *accountRepository) CreatePasswordResetToken(accountID int64, resetToken string) (int64, error) {
	// TODO: Disable all other tokens for this account

	result, err := a.DB.Exec(`INSERT
		INTO password_reset_tokens(
			account_id,
			reset_token
		) VALUES (
			?,
			?
		);
	`,
		accountID,
		resetToken,
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

// GetPasswordResetToken returns the password reset token with the given token.
func (a *accountRepository) GetPasswordResetToken(token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken

	err := a.DB.Get(&resetToken, `SELECT
		*
		FROM password_reset_tokens p
		WHERE p.reset_token = ?;
	`,
		token,
	)
	if err != nil {
		return &models.PasswordResetToken{}, err
	}

	return &resetToken, nil
}

// ResetPassword sets the password for the account associated with the reset token.
func (a *accountRepository) ResetPassword(token, password string) (bool, error) {
	tx, err := a.DB.Begin()
	_, err = tx.Exec(`UPDATE password_reset_tokens p
		INNER JOIN accounts a ON a.id = p.account_id
		SET is_used = 1
		WHERE reset_token = ?;
	`,
		token,
	)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec(`UPDATE accounts a
		INNER JOIN password_reset_tokens p ON p.account_id = a.id
		SET password = ?
		WHERE p.reset_token = ?;
	`,
		password,
		token,
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
