package services

import (
	"errors"

	"github.com/BrandonWade/enako/api/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/sirupsen/logrus"
)

var errCreatingAccount = errors.New("error creating account")

//go:generate counterfeiter -o fakes/fake_auth_service.go . AuthService
type AuthService interface {
	CreateAccount(username, email, password string) (int64, error)
}

type authService struct {
	logger *logrus.Logger
	repo   repositories.AuthRepository
}

func NewAuthService(logger *logrus.Logger, repo repositories.AuthRepository) AuthService {
	return &authService{
		logger,
		repo,
	}
}

func (a *authService) CreateAccount(username, email, password string) (int64, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

	return a.repo.CreateAccount(username, email, string(passwordHash))
}
