package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// PasswordResetRepository an interface for resetting passwords.
//go:generate counterfeiter -o fakes/fake_password_reset_repository.go . PasswordResetRepository
type PasswordResetRepository interface {
	CreatePasswordResetToken(accountID int64, resetToken string) (int64, error)
	GetPasswordResetTokenByPasswordResetToken(token string) (*models.PasswordResetToken, error)
	ResetPassword(token, password string) (bool, error)
}

type passwordResetRepository struct {
	DB *sqlx.DB
}

// NewPasswordResetRepository returns a new instance of an PasswordResetRepository.
func NewPasswordResetRepository(DB *sqlx.DB) PasswordResetRepository {
	return &passwordResetRepository{
		DB,
	}
}

// CreatePasswordResetToken creates an activation token for the given account ID.
func (a *passwordResetRepository) CreatePasswordResetToken(accountID int64, resetToken string) (int64, error) {
	tx, err := a.DB.Begin()
	result, err := a.DB.Exec(`INSERT
		INTO password_reset_tokens(
			account_id,
			reset_token,
			expires_at
		) VALUES (
			?,
			?,
			DATE_ADD(NOW(), INTERVAL 1 HOUR)
		);
	`,
		accountID,
		resetToken,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Disable all other pending tokens for this account
	_, err = tx.Exec(`UPDATE password_reset_tokens p
		SET p.status = 'disabled'
		WHERE p.account_id = ?
		AND p.status = 'pending'
		AND p.id != ?;
	`,
		accountID,
		id,
	)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetPasswordResetTokenByPasswordResetToken returns the password reset token with the given token.
func (a *passwordResetRepository) GetPasswordResetTokenByPasswordResetToken(token string) (*models.PasswordResetToken, error) {
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
func (a *passwordResetRepository) ResetPassword(token, password string) (bool, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return false, err
	}

	_, err = tx.Exec(`UPDATE password_reset_tokens p
		INNER JOIN accounts a ON a.id = p.account_id
		SET p.status = 'used'
		WHERE p.reset_token = ?
		AND p.status = 'pending';
	`,
		token,
	)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec(`UPDATE accounts a
		INNER JOIN password_reset_tokens p ON p.account_id = a.id
		SET a.password = ?
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
