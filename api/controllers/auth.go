package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"

	validator "gopkg.in/validator.v2"
)

var (
	ErrInvalidAccountPayload = errors.New("invalid account payload")
	ErrPasswordsDoNotMatch   = errors.New("passwords do not match")
	ErrCreatingAccount       = errors.New("error creating account")
	ErrFetchingSession       = errors.New("error fetching session")
)

//go:generate counterfeiter -o fakes/fake_auth_controller.go . AuthController
type AuthController interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.AuthService
}

func NewAuthController(logger *logrus.Logger, store helpers.CookieStorer, service services.AuthService) AuthController {
	return &authController{
		logger,
		store,
		service,
	}
}

func (a *authController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var userAccount models.UserAccount
	err := json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(ErrInvalidAccountPayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidAccountPayload))
		return
	}

	if err = validator.Validate(userAccount); err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(err))
		return
	}

	if userAccount.UserAccountPassword != userAccount.ConfirmPassword {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
		}).Error(ErrPasswordsDoNotMatch)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrPasswordsDoNotMatch))
		return
	}

	ID, err := a.service.CreateAccount(userAccount.UserAccountEmail, userAccount.UserAccountPassword)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
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

func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Login",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(ErrFetchingSession)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrFetchingSession))
		return
	}

	// TODO: Authenticate and return a real ID
	userAccountID := int64(123)

	session.Set("authenticated", true)
	session.Set("user_account_id", userAccountID)
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func (a *authController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Logout",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(ErrFetchingSession)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrFetchingSession))
		return
	}

	session.Set("authenticated", false)
	session.Set("user_account_id", int64(0))
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
