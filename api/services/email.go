package services

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// EmailService an interface for sending emails.
//go:generate counterfeiter -o fakes/fake_email_service.go . EmailService
type EmailService interface {
	SendActivateAccountEmail(token, email string) error
}

type emailService struct {
	logger *logrus.Logger
}

// NewEmailService returns a new instance of an EmailService.
func NewEmailService(logger *logrus.Logger) EmailService {
	return &emailService{
		logger,
	}
}

// SendActivateAccountEmail sends an email with an activation link to the provided email.
func (e *emailService) SendActivateAccountEmail(email, token string) error {
	link := fmt.Sprintf("%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token)

	// TODO: Implement
	fmt.Println(link)

	return nil
}
