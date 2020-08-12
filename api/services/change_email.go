package services

import (
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

const (
	// ChangeEmailTokenLength the length of a change email token.
	ChangeEmailTokenLength = 64
)

// ChangeEmailService an interface for changing the email address associated with an account.
//go:generate counterfeiter -o fakes/fake_change_email_service.go . ChangeEmailService
type ChangeEmailService interface {
	RequestEmailChange(accountID int64) (string, error)
}

type changeEmailService struct {
	logger       *logrus.Logger
	generator    helpers.TokenGenerator
	emailService EmailService
	repo         repositories.ChangeEmailRepository
	accountRepo  repositories.AccountRepository
}

// NewChangeEmailService returns a new instance of an ChangeEmailService.
func NewChangeEmailService(logger *logrus.Logger, generator helpers.TokenGenerator, emailService EmailService, repo repositories.ChangeEmailRepository, accountRepo repositories.AccountRepository) ChangeEmailService {
	return &changeEmailService{
		logger,
		generator,
		emailService,
		repo,
		accountRepo,
	}
}

// RequestEmailChange requests an email change for the account with the given ID.
func (c *changeEmailService) RequestEmailChange(accountID int64) (string, error) {
	account, err := c.accountRepo.GetAccountByID(accountID)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method":    "ChangeEmailService.RequestEmailChange",
			"accountID": accountID,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingEmailChange()
	}

	token := c.generator.CreateToken(ChangeEmailTokenLength)
	_, err = c.repo.CreateChangeEmailToken(accountID, token)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method":    "ChangeEmailService.RequestEmailChange",
			"accountID": accountID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingEmailChange()
	}

	err = c.emailService.SentChangeEmailEmail(account.Email, token)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method":    "ChangeEmailService.RequestEmailChange",
			"accountID": accountID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingEmailChange()
	}

	return account.Email, nil
}
