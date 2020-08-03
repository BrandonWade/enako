package services

import (
	"time"

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
	CreateAccount(email, password string) (int64, error)
	RegisterUser(email, password string) (int64, error)
	VerifyAccount(email, password string) (int64, error)
	ActivateAccount(token string) (bool, error)
	ChangePassword(accountID int64, currentPassword, password string) (int64, error)
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

// CreateAccount creates an account with the given email and password.
func (a *accountService) CreateAccount(email, password string) (int64, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AccountService.CreateAccount",
			"password": password,
		}).Error(err.Error())
		return 0, err
	}

	id, err := a.repo.CreateAccount(email, string(hash))
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountService.CreateAccount",
			"email":  email,
		}).Error(err.Error())
		return 0, err
	}

	return id, nil
}

// RegisterUser creates an account, generates an activation token, and sends an activation email.
func (a *accountService) RegisterUser(email, password string) (int64, error) {
	accountID, err := a.CreateAccount(email, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountService.RegisterUser",
			"email":  email,
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

// VerifyAccount checks whether or not an account exists for the given email and password.
func (a *accountService) VerifyAccount(email, password string) (int64, error) {
	account, err := a.repo.GetAccountByEmail(email)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountService.VerifyAccount",
			"email":  email,
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
		activationToken, err := a.repo.GetActivationTokenByAccountID(account.ID)
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"method":    "AccountService.VerifyAccount",
				"accountID": account.ID,
			}).Error(err.Error())
			return 0, err
		}

		lastSent, err := time.Parse("2006-01-02 15:04:05", activationToken.LastSentAt)
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"method":  "AccountService.VerifyAccount",
				"tokenID": activationToken.ID,
			}).Error(err.Error())
			return 0, err
		}

		if time.Now().After(lastSent.Add(1 * time.Hour)) {
			err = a.emailService.SendAccountActivationEmail(account.Email, activationToken.ActivationToken)
			if err != nil {
				a.logger.WithFields(logrus.Fields{
					"method":  "AccountService.VerifyAccount",
					"tokenID": activationToken.ID,
				}).Error(err.Error())
				return 0, err
			}

			_, err := a.repo.UpdateActivationTokenLastSentAt(activationToken.ID)
			if err != nil {
				a.logger.WithFields(logrus.Fields{
					"method":  "AccountService.VerifyAccount",
					"tokenID": activationToken.ID,
				}).Error(err.Error())
				return 0, err
			}

			return 0, helpers.ErrorActivationEmailResent()
		}

		a.logger.WithFields(logrus.Fields{
			"method": "AccountService.VerifyAccount",
			"email":  email,
		}).Info(helpers.ErrorAccountNotActivated())
		return 0, helpers.ErrorAccountNotActivated()
	}

	return account.ID, nil
}

// ActivateAccount activates the account with the given token.
func (a *accountService) ActivateAccount(token string) (bool, error) {
	return a.repo.ActivateAccount(token)
}

// ChangePassword updates the password for the account with the given id.
func (a *accountService) ChangePassword(accountID int64, currentPassword, password string) (int64, error) {
	account, err := a.repo.GetAccountByID(accountID)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.ChangePassword",
			"accountID": accountID,
		}).Error(err.Error())
		return 0, err
	}

	err = a.hasher.Compare(account.Password, currentPassword)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.ChangePassword",
			"accountID": accountID,
		}).Error(helpers.ErrorPasswordsDoNotMatch())
		return 0, err
	}

	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.ChangePassword",
			"accountID": accountID,
		}).Error(err.Error())
		return 0, err
	}

	count, err := a.repo.ChangePassword(accountID, hash)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.ChangePassword",
			"accountID": accountID,
		}).Error(err.Error())
		return 0, err
	}

	err = a.emailService.SendPasswordUpdatedEmail(account.Email)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "AccountService.ChangePassword",
			"accountID": accountID,
		}).Error(err.Error())
		return 0, err
	}

	return count, nil
}
