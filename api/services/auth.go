package services

import (
	"github.com/BrandonWade/enako/api/repositories"
	"golang.org/x/crypto/bcrypt"
)

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
		// TODO: Handle
	}

	return a.repo.CreateAccount(email, string(passwordHash))
}
