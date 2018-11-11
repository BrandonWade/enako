package services

import "github.com/BrandonWade/enako/api/repositories"

type AuthService interface {
	CreateAccount()
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{
		repo,
	}
}

func (a *authService) CreateAccount() {

}
