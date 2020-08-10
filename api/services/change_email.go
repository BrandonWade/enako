package services

import (
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

// ChangeEmailService an interface for changing the email address associated with an account.
//go:generate counterfeiter -o fakes/fake_change_email_service.go . ChangeEmailService
type ChangeEmailService interface {
	RequestEmailChange(email string) (string, error)
}

type changeEmailService struct {
	logger *logrus.Logger
	repo   repositories.ChangeEmailRepository
}

// NewChangeEmailService returns a new instance of an ChangeEmailService.
func NewChangeEmailService(logger *logrus.Logger, repo repositories.ChangeEmailRepository) ChangeEmailService {
	return &changeEmailService{
		logger,
		repo,
	}
}

// RequestEmailChange
func (c *changeEmailService) RequestEmailChange(email string) (string, error) {
	// TODO: Implement

	return "", nil
}
