package services

import (
	"database/sql"
	"errors"

	"github.com/BrandonWade/enako/api/models"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"

	"github.com/sirupsen/logrus"
)

const (
	// ActivationTokenLength the length of an account activation token.
	ActivationTokenLength = 64

	// PasswordResetTokenLength the length of a password reset token.
	PasswordResetTokenLength = 64
)

// AccountService an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_account_service.go . AccountService
type AccountService interface {
	CreateAccount(username, email, password string) (int64, error)
	RegisterUser(username, email, password string) (int64, error)
	VerifyAccount(username, password string) (int64, error)
	ActivateAccount(token string) (bool, error)
	RequestPasswordReset(username string) (string, error)
	GetPasswordResetToken(token string) (*models.PasswordResetToken, error)
	ResetPassword(token, password string) (bool, error)
	NotifyOfPasswordReset(token string) error
}

type accountService struct {
	logger       *logrus.Logger
	hasher       helpers.PasswordHasher
	obfuscator   helpers.EmailObfuscator
	generator    helpers.TokenGenerator
	emailService EmailService
	repo         repositories.AccountRepository
}

// NewAccountService returns a new instance of an AccountService.
func NewAccountService(logger *logrus.Logger, hasher helpers.PasswordHasher, obfuscator helpers.EmailObfuscator, generator helpers.TokenGenerator, emailService EmailService, repo repositories.AccountRepository) AccountService {
	return &accountService{
		logger,
		hasher,
		obfuscator,
		generator,
		emailService,
		repo,
	}
}

// CreateAccount creates an account with the given username, email, and password.
func (a *accountService) CreateAccount(username, email, password string) (int64, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.CreateAccount",
			"password": password,
		}).Error(err.Error())
		return 0, err
	}

	id, err := a.repo.CreateAccount(username, email, string(hash))
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.CreateAccount",
			"username": username,
		}).Error(err.Error())
		return 0, err
	}

	return id, nil
}

// RegisterUser creates an account, generates an activation token, and sends an activation email.
func (a *accountService) RegisterUser(username, email, password string) (int64, error) {
	accountID, err := a.CreateAccount(username, email, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.RegisterUser",
			"username": username,
			"email":    email,
		}).Error(err.Error())
		return 0, err
	}

	token := a.generator.CreateToken(ActivationTokenLength)
	_, err = a.repo.CreateActivationToken(accountID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.RegisterUser",
			"accountID": accountID,
			"token":     token,
		}).Error(err.Error())
		return 0, err
	}

	err = a.emailService.SendAccountActivationEmail(email, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.RegisterUser",
			"email":     email,
			"accountID": accountID,
			"token":     token,
		}).Error(err.Error())
		return 0, err
	}

	return accountID, nil
}

// VerifyAccount checks whether or not an account exists for the given username and password.
func (a *accountService) VerifyAccount(username, password string) (int64, error) {
	account, err := a.repo.GetAccount(username)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.VerifyAccount",
			"username": username,
		}).Error(err.Error())
		return 0, err
	}

	err = a.hasher.Compare(account.Password, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.VerifyAccount",
			"password": password,
		}).Error(err.Error())
		return 0, err
	}

	if !account.IsActivated {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.VerifyAccount",
			"username": username,
		}).Info(helpers.ErrorAccountNotActivated())
		return 0, helpers.ErrorAccountNotActivated()
	}

	return account.ID, nil
}

// ActivateAccount activates the account with the given token.
func (a *accountService) ActivateAccount(token string) (bool, error) {
	return a.repo.ActivateAccount(token)
}

// RequestPasswordReset requests a password reset for the account with the given username.
func (a *accountService) RequestPasswordReset(username string) (string, error) {
	account, err := a.repo.GetAccountByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			a.logger.WithFields(logrus.Fields{
				"method":   "AccountService.RequestPasswordReset",
				"username": username,
			}).Info(err.Error())

			return "", helpers.ErrorAccountNotFound()
		}

		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.RequestPasswordReset",
			"username": username,
		}).Error(err.Error())

		return "", helpers.ErrorRequestingPasswordReset()
	}

	token := a.generator.CreateToken(PasswordResetTokenLength)
	_, err = a.repo.CreatePasswordResetToken(account.ID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.RequestPasswordReset",
			"accountID": account.ID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	err = a.emailService.SendPasswordResetEmail(account.Email, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.RequestPasswordReset",
			"username": username,
			"email":    account.Email,
			"token":    token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	email, err := a.obfuscator.Obfuscate(account.Email)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.RequestPasswordReset",
			"username": username,
			"email":    account.Email,
		}).Error(err.Error())

		return "", helpers.ErrorRequestingPasswordReset()
	}

	return email, nil
}

// GetPasswordResetToken returns the password reset token with the given token.
func (a *accountService) GetPasswordResetToken(token string) (*models.PasswordResetToken, error) {
	return a.repo.GetPasswordResetToken(token)
}

// ResetPassword sets the password for the account associated with the reset token.
func (a *accountService) ResetPassword(token, password string) (bool, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.ResetPassword",
			"password": password,
		}).Error(err.Error())
		return false, err
	}

	return a.repo.ResetPassword(token, string(hash))
}

// NotifyOfPasswordReset notifies the account owner that their password was reset.
func (a *accountService) NotifyOfPasswordReset(token string) error {
	account, err := a.repo.GetAccountByPasswordResetToken(token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountService.NotifyOfPasswordReset",
			"token":  token,
		}).Error(err.Error())
		return err
	}

	return a.emailService.SendPasswordUpdatedEmail(account.Email)
}
