package services

import (
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/dchest/uniuri"

	"github.com/sirupsen/logrus"
)

const (
	// ActivationTokenLength the length of an account activation token.
	ActivationTokenLength = 64
)

// AuthService an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_auth_service.go . AuthService
type AuthService interface {
	CreateAccount(username, email, password string) (int64, string, error)
	VerifyAccount(username, password string) (int64, error)
	ActivateAccount(token string) (bool, error)
}

type authService struct {
	logger *logrus.Logger
	hasher helpers.PasswordHasher
	repo   repositories.AuthRepository
}

// NewAuthService returns a new instance of an AuthService.
func NewAuthService(logger *logrus.Logger, hasher helpers.PasswordHasher, repo repositories.AuthRepository) AuthService {
	return &authService{
		logger,
		hasher,
		repo,
	}
}

// CreateAccount creates an account with the given username, email, and password.
func (a *authService) CreateAccount(username, email, password string) (int64, string, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.CreateAccount",
			"password": password,
		}).Error(err.Error())
		return 0, "", err
	}

	id, err := a.repo.CreateAccount(username, email, string(hash))
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.CreateAccount",
			"username": username,
		}).Error(err.Error())
		return 0, "", err
	}

	token := uniuri.NewLen(ActivationTokenLength)
	_, err = a.repo.CreateActivationToken(id, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.CreateAccount",
			"username": username,
			"token":    token,
		}).Error(err.Error())
		return 0, "", err
	}

	return id, token, nil
}

// VerifyAccount checks whether or not an account exists for the given username and password.
func (a *authService) VerifyAccount(username, password string) (int64, error) {
	account, err := a.repo.GetAccount(username)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.VerifyAccount",
			"username": username,
		}).Error(err.Error())
		return 0, err
	}

	err = a.hasher.Compare(account.Password, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.VerifyAccount",
			"password": password,
		}).Error(err.Error())
		return 0, err
	}

	if !account.IsActivated {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.VerifyAccount",
			"username": username,
		}).Info(helpers.ErrorAccountNotActivated())
		return 0, helpers.ErrorAccountNotActivated()
	}

	return account.ID, nil
}

// ActivateAccount activates the account with the given token.
func (a *authService) ActivateAccount(token string) (bool, error) {
	return a.repo.ActivateAccount(token)
}
