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
			"password": password,
		}).Error(err.Error())
		return 0, err
	}

	return a.repo.CreateAccount(username, email, string(hash))
}

// VerifyAccount ...
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

	return account.ID, nil
}
