package repositories

import (
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// ChangeEmailRepository an interface for changing the email address associated with an account..
//go:generate counterfeiter -o fakes/fake_change_email_repository.go . ChangeEmailRepository
type ChangeEmailRepository interface {
	// TODO: Implement
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
