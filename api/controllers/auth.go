package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"

	log "github.com/sirupsen/logrus"
	validator "gopkg.in/validator.v2"
)

var (
	ErrInvalidAccountPayload = errors.New("invalid account payload")
	ErrPasswordsDoNotMatch   = errors.New("passwords do not match")
	ErrCreatingAccount       = errors.New("error creating account")
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
		}).Error(ErrInvalidAccountPayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidAccountPayload))
		return
	}

	if err = validator.Validate(userAccount); err != nil {
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(err))
		return
	}

	if userAccount.UserAccountPassword != userAccount.ConfirmPassword {
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
		}).Error(ErrPasswordsDoNotMatch)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrPasswordsDoNotMatch))
		return
	}

	ID, err := a.service.CreateAccount(userAccount.UserAccountEmail, userAccount.UserAccountPassword)
	if err != nil {
		log.WithFields(log.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(ErrCreatingAccount)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrCreatingAccount))
		return
	}

	userAccount.ID = ID
	userAccount.UserAccountPassword = ""
	userAccount.ConfirmPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userAccount)
	return
}
