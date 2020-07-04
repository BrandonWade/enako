package services

import (
	"database/sql"
	"errors"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/dchest/uniuri"

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
	CreateActivationToken(accountID int64, token string) (int64, error)
	VerifyAccount(username, password string) (int64, error)
	ActivateAccount(token string) (bool, error)
	GetAccountByUsername(username string) (*models.Account, error)
	CreatePasswordResetToken(accountID int64, token string) (int64, error)
	RequestPasswordReset(username string) (string, error)
}

type accountService struct {
	logger       *logrus.Logger
	hasher       helpers.PasswordHasher
	obfuscator   helpers.EmailObfuscator
	emailService EmailService
	repo         repositories.AccountRepository
}

// NewAccountService returns a new instance of an AccountService.
func NewAccountService(logger *logrus.Logger, hasher helpers.PasswordHasher, obfuscator helpers.EmailObfuscator, emailService EmailService, repo repositories.AccountRepository) AccountService {
	return &accountService{
		logger,
		hasher,
		obfuscator,
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

// CreateActivationToken registers an activation token for the account with the given id.
func (a *accountService) CreateActivationToken(accountID int64, token string) (int64, error) {
	id, err := a.repo.CreateActivationToken(accountID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.CreateActivationToken",
			"accountID": accountID,
			"token":     token,
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

	token := uniuri.NewLen(ActivationTokenLength)
	_, err = a.CreateActivationToken(accountID, token)
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

// GetAccountByUsername returns the account with the given username.
func (a *accountService) GetAccountByUsername(username string) (*models.Account, error) {
	return a.repo.GetAccountByUsername(username)
}

// CreatePasswordResetToken registers a password reset token for the account with the given id.
func (a *accountService) CreatePasswordResetToken(accountID int64, token string) (int64, error) {
	id, err := a.repo.CreatePasswordResetToken(accountID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.CreatePasswordResetToken",
			"accountID": accountID,
			"token":     token,
		}).Error(err.Error())
		return 0, err
	}

	return id, nil
}

// RequestPasswordReset requests a password reset for the account with the given username.
func (a *accountService) RequestPasswordReset(username string) (string, error) {
	account, err := a.GetAccountByUsername(username)
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

	token := uniuri.NewLen(PasswordResetTokenLength)
	_, err = a.CreatePasswordResetToken(account.ID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.RequestPasswordReset",
			"accountID": account.ID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	err = a.emailService.SendAccountActivationEmail(account.Email, token)
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
