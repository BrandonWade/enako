package services

import (
	"fmt"
	"os"

	"github.com/BrandonWade/enako/api/clients"
	"github.com/sirupsen/logrus"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// EmailService an interface for sending emails.
//go:generate counterfeiter -o fakes/fake_email_service.go . EmailService
type EmailService interface {
	SendAccountActivationEmail(token, email string) error
	SendPasswordResetEmail(email, token string) error
}

type emailService struct {
	logger          *logrus.Logger
	templateService TemplateService
	emailClient     clients.MailjetClient
}

// NewEmailService returns a new instance of an EmailService.
func NewEmailService(logger *logrus.Logger, templateService TemplateService, emailClient clients.MailjetClient) EmailService {
	return &emailService{
		logger,
		templateService,
		emailClient,
	}
}

// SendAccountActivationEmail sends an email with an activation link to the provided email.
func (e *emailService) SendAccountActivationEmail(email, token string) error {
	link := fmt.Sprintf("%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token)

	template, err := e.templateService.GenerateAccountActivationEmail(link)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendAccountActivationEmail",
			"email":  email,
			"token":  token,
		}).Error(err.Error())
		return err
	}

	message := mailjet.InfoMessagesV31{
		From: &mailjet.RecipientV31{
			Email: fmt.Sprintf("register@%s", os.Getenv("ENAKO_DOMAIN")),
			Name:  "Enako",
		},
		To: &mailjet.RecipientsV31{
			mailjet.RecipientV31{
				Email: email,
			},
		},
		Subject:  "Complete Registration",
		HTMLPart: template,
		CustomID: "EnakoAccountActivationEmail",
	}

	err = e.emailClient.Send(message)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendAccountActivationEmail",
			"email":  email,
			"token":  token,
		}).Error(err.Error())
		return err
	}

	return nil
}

// SendPasswordResetEmail sends an email with a password reset link to the provided email.
func (e *emailService) SendPasswordResetEmail(email, token string) error {
	link := fmt.Sprintf("%s/api/v1/accounts/password/reset?t=%s", os.Getenv("API_HOST"), token)

	template, err := e.templateService.GeneratePasswordResetEmail(link)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendPasswordResetEmail",
			"email":  email,
			"token":  token,
		}).Error(err.Error())
		return err
	}

	message := mailjet.InfoMessagesV31{
		From: &mailjet.RecipientV31{
			Email: fmt.Sprintf("accounts@%s", os.Getenv("ENAKO_DOMAIN")),
			Name:  "Enako",
		},
		To: &mailjet.RecipientsV31{
			mailjet.RecipientV31{
				Email: email,
			},
		},
		Subject:  "Password Reset",
		HTMLPart: template,
		CustomID: "EnakoPasswordResetEmail",
	}

	err = e.emailClient.Send(message)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendPasswordResetEmail",
			"email":  email,
			"token":  token,
		}).Error(err.Error())
		return err
	}

	return nil
}
