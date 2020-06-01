package services

import (
	"errors"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"

	"github.com/sirupsen/logrus"
)

var errCreatingAccount = errors.New("error creating account")
var errVerifyingAccount = errors.New("error verifying account")

//go:generate counterfeiter -o fakes/fake_auth_service.go . AuthService
type AuthService interface {
	CreateAccount(username, email, password string) (int64, error)
	VerifyAccount(username, password string) (int64, error)
}

type authService struct {
	logger *logrus.Logger
	hasher helpers.PasswordHasher
	repo   repositories.AuthRepository
}

func NewAuthService(logger *logrus.Logger, hasher helpers.PasswordHasher, repo repositories.AuthRepository) AuthService {
	return &authService{
		logger,
		hasher,
		repo,
	}
}

func (a *authService) CreateAccount(username, email, password string) (int64, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.CreateAccount",
			"username": username,
			"email":    email,
			"password": password,
			"err":      err.Error(),
		}).Error(errCreatingAccount)

		return 0, errCreatingAccount
	}

	return a.repo.CreateAccount(username, email, string(hash))
}

func (a *authService) VerifyAccount(username, password string) (int64, error) {
	account, err := a.repo.GetAccount(username)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.VerifyingAccount",
			"username": username,
			"password": password,
			"err":      err.Error(),
		}).Error(errVerifyingAccount)

		return 0, errVerifyingAccount
	}

	err = a.hasher.Compare(account.Password, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.VerifyingAccount",
			"username": username,
			"password": password,
			"err":      err.Error(),
		}).Error(errVerifyingAccount)

		return 0, errVerifyingAccount
	}

	return account.ID, nil
}
