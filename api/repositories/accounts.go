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
	GetAccount(email string) (*models.Account, error)
	CreateAccount(email, password string) (int64, error)
	CreateActivationToken(accountID int64, activationToken string) (int64, error)
	ActivateAccount(token string) (bool, error)
	GetAccountByEmail(email string) (*models.Account, error)
	GetAccountByPasswordResetToken(token string) (*models.Account, error)

	GetActivationTokenByAccountID(accountID int64) (*models.ActivationToken, error)
	UpdateActivationTokenLastSentAt(tokenID int64) (int64, error)
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

// GetAccount returns an account with the given email.
func (a *accountRepository) GetAccount(email string) (*models.Account, error) {
	account := models.Account{}

	err := a.DB.Get(&account, `SELECT
		*
		FROM accounts a
		WHERE a.email = ?;
	`,
		email,
	)
	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil
}

// CreateAccount creates an account with the given email, email, and password.
func (a *accountRepository) CreateAccount(email, password string) (int64, error) {
	result, err := a.DB.Exec(`INSERT
		INTO accounts(
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

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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

// GetAccountByEmail returns the account with the given email.
func (a *accountRepository) GetAccountByEmail(email string) (*models.Account, error) {
	var account models.Account

	err := a.DB.Get(&account, `SELECT
		*
		FROM accounts a
		WHERE a.email = ?;
	`,
		email,
	)
	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil
}

// GetAccountByPasswordResetToken returns the account associated with the given password reset token.
func (a *accountRepository) GetAccountByPasswordResetToken(token string) (*models.Account, error) {
	var account models.Account

	err := a.DB.Get(&account, `SELECT
		a.*
		FROM accounts a
		INNER JOIN password_reset_tokens p ON p.account_id = a.id
		WHERE p.reset_token = ?;
	`,
		token,
	)
	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil
}

// GetActivationTokenByAccountID returns the activation token associated with the given account id.
func (a *accountRepository) GetActivationTokenByAccountID(accountID int64) (*models.ActivationToken, error) {
	var activationToken models.ActivationToken

	err := a.DB.Get(&activationToken, `SELECT
		t.*
		FROM account_activation_tokens t
		WHERE t.account_id = ?;
	`,
		accountID,
	)
	if err != nil {
		return &models.ActivationToken{}, err
	}

	return &activationToken, nil
}

// UpdateActivationTokenLastSentAt sets the last sent at timestamp for the activation token with the given id.
func (a *accountRepository) UpdateActivationTokenLastSentAt(tokenID int64) (int64, error) {
	result, err := a.DB.Exec(`UPDATE account_activation_tokens
		SET last_sent_at = NOW()
		WHERE id = ?;
	`,
		tokenID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
