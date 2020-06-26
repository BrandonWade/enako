package services

import (
	"fmt"

	"github.com/BrandonWade/enako/api/clients"
	"github.com/sirupsen/logrus"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// EmailService an interface for sending emails.
//go:generate counterfeiter -o fakes/fake_email_service.go . EmailService
type EmailService interface {
	SendActivateAccountEmail(token, email string) error
}

type emailService struct {
	logger *logrus.Logger
	client clients.MailjetClient
}

// NewEmailService returns a new instance of an EmailService.
func NewEmailService(logger *logrus.Logger, client clients.MailjetClient) EmailService {
	return &emailService{
		logger,
		client,
	}
}

// SendActivateAccountEmail sends an email with an activation link to the provided email.
func (e *emailService) SendActivateAccountEmail(email, token string) error {
	// link := fmt.Sprintf("%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token)

	message := mailjet.InfoMessagesV31{
		From: &mailjet.RecipientV31{
			Email: "",
			Name:  "",
		},
		To: &mailjet.RecipientsV31{
			mailjet.RecipientV31{
				Email: "",
				Name:  "",
			},
		},
		Subject:  "Sample Email 123",
		TextPart: "My first Mailjet email",
		HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
		CustomID: "AppGettingStartedTest",
	}

	err := e.client.Send(message)
	if err != nil {
		fmt.Printf("ERR SENDING ACTIVATE EMAIL: %s", err.Error())
		return err
	}

	return nil
}
