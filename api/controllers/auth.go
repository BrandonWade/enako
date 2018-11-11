package controllers

import (
	"net/http"

	"github.com/BrandonWade/enako/api/services"
)

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
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	// TODO: Validate inputs

	if password != confirmPassword {
		// TODO: Handle
	}

	a.service.CreateAccount(email, password)
}
