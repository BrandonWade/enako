package services

import (
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
	// TODO: Implement

	// link := fmt.Sprintf("%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token)

	// auth := smtp.PlainAuth("", "register@enako.ca", "password", "mail.example.com")
	// to := []string{email}
	// msg := []byte("To: recipient@example.net\r\n" +
	// 	"Subject: Complete Account Registration\r\n" +
	// 	"\r\n" +
	// 	"Activation link: " + link)

	// err := smtp.SendMail("mail.example.com:25", auth, "register@enako.ca", to, msg)
	// if err != nil {
	// 	e.logger.WithFields(logrus.Fields{
	// 		"method": "EmailService.SendActivateAccountEmail",
	// 		"email":  email,
	// 		"token":  token,
	// 	}).Error(err.Error())
	// 	return err
	// }

	return nil
}
