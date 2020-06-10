package services

import (
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"

	"github.com/sirupsen/logrus"
)

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

// NewAuthService ...
func NewAuthService(logger *logrus.Logger, hasher helpers.PasswordHasher, repo repositories.AuthRepository) AuthService {
	return &authService{
		logger,
		hasher,
		repo,
	}
}

// CreateAccount ...
func (a *authService) CreateAccount(username, email, password string) (int64, error) {
	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthService.CreateAccount",
			"username": username,
			"email":    email,
			"password": password,
			"err":      err.Error(),
		}).Error(helpers.ErrorCreatingAccount())

		return 0, helpers.ErrorCreatingAccount()
	}

	return a.repo.CreateAccount(username, email, string(hash))
}

// VerifyAccount ...
func (a *authService) VerifyAccount(username, password string) (int64, error) {
	a.logger.WithFields(logrus.Fields{
		"method":   "AuthService.VerifyAccount",
		"username": username,
		"password": password,
	})

	account, err := a.repo.GetAccount(username)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(helpers.ErrorVerifyingAccount())

		return 0, helpers.ErrorVerifyingAccount()
	}

	err = a.hasher.Compare(account.Password, password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(helpers.ErrorVerifyingAccount())

		return 0, helpers.ErrorVerifyingAccount()
	}

	return account.ID, nil
}
