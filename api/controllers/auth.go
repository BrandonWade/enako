package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
)

//go:generate counterfeiter -o fakes/fake_auth_controller.go . AuthController
type AuthController interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{
		service,
	}
}

func (a *authController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var userAccount models.UserAccount
	err := json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		// TODO: Handle
	}

	// TODO: Validate inputs

	if userAccount.UserAccountPassword != userAccount.ConfirmPassword {
		// TODO: Handle
	}

	ID, err := a.service.CreateAccount(userAccount.UserAccountEmail, userAccount.UserAccountPassword)
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(ID)
}
