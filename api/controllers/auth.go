package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"

	log "github.com/sirupsen/logrus"
)

var (
	errInvalidAccountPayload = errors.New("invalid account payload")
	errPasswordsDoNotMatch   = errors.New("passwords do not match")
	errCreatingAccount       = errors.New("error creating account")
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
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(errInvalidAccountPayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(errInvalidAccountPayload))
		return
	}

	// TODO: Validate inputs

	if userAccount.UserAccountPassword != userAccount.ConfirmPassword {
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(errPasswordsDoNotMatch)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(errPasswordsDoNotMatch))
		return
	}

	ID, err := a.service.CreateAccount(userAccount.UserAccountEmail, userAccount.UserAccountPassword)
	if err != nil {
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(errCreatingAccount)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errCreatingAccount))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ID)
	return
}
