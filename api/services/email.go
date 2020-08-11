package services

import (
	"fmt"
	"os"

	"github.com/BrandonWade/enako/api/clients"
	"github.com/sirupsen/logrus"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

const (
	accountActivationEmailID = "EnakoAccountActivationEmail"
	passwordResetEmailID     = "EnakoPasswordResetEmail"
	passwordUpdatedEmailID   = "EnakoPasswordUpdatedEmail"
)

// EmailService an interface for sending emails.
//go:generate counterfeiter -o fakes/fake_email_service.go . EmailService
type EmailService interface {
	SendAccountActivationEmail(email, token string) error
	SendPasswordResetEmail(email, token string) error
	SendPasswordUpdatedEmail(email string) error
	SentChangeEmailEmail(email, token string) error
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
	link := fmt.Sprintf("http://%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token)

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
		CustomID: accountActivationEmailID,
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
	link := fmt.Sprintf("http://%s/api/v1/accounts/password/reset?t=%s", os.Getenv("API_HOST"), token)

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
		CustomID: passwordResetEmailID,
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

// SendPasswordUpdatedEmail sends an email when a password has been updated.
func (e *emailService) SendPasswordUpdatedEmail(email string) error {
	link := fmt.Sprintf("http://%s/password", os.Getenv("API_HOST"))

	template, err := e.templateService.GeneratePasswordUpdatedEmail(link)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendPasswordUpdatedEmail",
			"email":  email,
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
		Subject:  "Password Updated",
		HTMLPart: template,
		CustomID: passwordUpdatedEmailID,
	}

	err = e.emailClient.Send(message)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailService.SendPasswordUpdatedEmail",
			"email":  email,
		}).Error(err.Error())
		return err
	}

	return nil
}

// SentChangeEmailEmail sends an email with a change email link to the provided email.
func (e *emailService) SentChangeEmailEmail(email, token string) error {
	// TODO: Implement

	return nil
}
