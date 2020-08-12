package repositories

import (
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// ChangeEmailRepository an interface for changing the email address associated with an account..
//go:generate counterfeiter -o fakes/fake_change_email_repository.go . ChangeEmailRepository
type ChangeEmailRepository interface {
	CreateChangeEmailToken(accountID int64, token string) (int64, error)
}

type changeEmailRepository struct {
	DB *sqlx.DB
}

// NewChangeEmailRepository returns a new instance of an ChangeEmailRepository.
func NewChangeEmailRepository(DB *sqlx.DB) ChangeEmailRepository {
	return &changeEmailRepository{
		DB,
	}
}

// CreateChangeEmailToken creates a change email token for the given account ID.
func (c *changeEmailRepository) CreateChangeEmailToken(accountID int64, token string) (int64, error) {
	tx, err := c.DB.Begin()
	result, err := c.DB.Exec(`INSERT
		INTO change_email_tokens(
			account_id,
			change_token,
			expires_at
		) VALUES (
			?,
			?,
			DATE_ADD(NOW(), INTERVAL 1 HOUR)
		);
	`,
		accountID,
		token,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Disable all other pending tokens for this account
	_, err = tx.Exec(`UPDATE change_email_tokens p
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
