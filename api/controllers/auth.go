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
	ErrInvalidAccountPayload     = errors.New("invalid account payload")
	ErrPasswordsDoNotMatch       = errors.New("passwords do not match")
	ErrCreatingAccount           = errors.New("error creating account")
	ErrFetchingSession           = errors.New("error fetching session")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
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

	if userAccount.Password != userAccount.ConfirmPassword {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.CreateAccount",
			"ip":     r.RemoteAddr,
		}).Error(ErrPasswordsDoNotMatch)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrPasswordsDoNotMatch))
		return
	}

	ID, err := a.service.CreateAccount(userAccount.Username, userAccount.Email, userAccount.Password)
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
	userAccount.Password = ""
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

	var userAccount models.UserAccount
	err = json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Login",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(ErrInvalidAccountPayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidAccountPayload))
		return
	}

	ID, err := a.service.VerifyAccount(userAccount.Username, userAccount.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthController.Login",
			"ip":       r.RemoteAddr,
			"username": userAccount.Username,
			"err":      err.Error(),
		}).Error(ErrInvalidUsernameOrPassword)

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidUsernameOrPassword))
		return
	}

	session.Set("authenticated", true)
	session.Set("user_account_id", ID)
	session.Save(r, w)

	userAccount.ID = ID
	userAccount.Password = ""

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userAccount)
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

	session.Delete()
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	return
}
