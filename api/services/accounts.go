package services

import (
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"

	"github.com/sirupsen/logrus"
)

const (
	// ActivationTokenLength the length of an account activation token.
	ActivationTokenLength = 64
)

// AccountService an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_account_service.go . AccountService
type AccountService interface {
	CreateAccount(username, email, password string) (int64, error)
	RegisterUser(username, email, password string) (int64, error)
	VerifyAccount(username, password string) (int64, error)
	ActivateAccount(token string) (bool, error)
}

type accountService struct {
	logger       *logrus.Logger
	hasher       helpers.PasswordHasher
	generator    helpers.TokenGenerator
	emailService EmailService
	repo         repositories.AccountRepository
}

// NewAccountService returns a new instance of an AccountService.
func NewAccountService(logger *logrus.Logger, hasher helpers.PasswordHasher, generator helpers.TokenGenerator, emailService EmailService, repo repositories.AccountRepository) AccountService {
	return &accountService{
		logger,
		hasher,
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
