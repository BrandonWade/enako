package services

import (
	"errors"

	"github.com/BrandonWade/enako/api/repositories"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

var errCreatingAccount = errors.New("error creating account")

//go:generate counterfeiter -o fakes/fake_auth_service.go . AuthService
type AuthService interface {
	CreateAccount(email, password string) (int64, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{
		repo,
	}
}

func (a *authService) CreateAccount(email, password string) (int64, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"method":   "AuthService.CreateAccount",
			"email":    email,
			"password": password,
			"err":      err.Error(),
		}).Error(errCreatingAccount)

		return 0, errCreatingAccount
	}

	return a.repo.CreateAccount(email, string(passwordHash))
}
