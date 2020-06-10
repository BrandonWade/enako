package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
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

// NewAuthController ...
func NewAuthController(logger *logrus.Logger, store helpers.CookieStorer, service services.AuthService) AuthController {
	return &authController{
		logger,
		store,
		service,
	}
}

// CreateAccount ...
func (a *authController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	a.logger.WithFields(logrus.Fields{
		"method": "AuthController.CreateAccount",
		"ip":     r.RemoteAddr,
	})

	userAccount, ok := r.Context().Value(middleware.ContextUserAccountKey).(models.UserAccount)
	if !ok {
		a.logger.Error(helpers.ErrorRetrievingAccount())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	ID, err := a.service.CreateAccount(userAccount.Username, userAccount.Email, userAccount.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(helpers.ErrorCreatingAccount())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	userAccount.ID = ID
	userAccount.Password = ""
	userAccount.ConfirmPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userAccount)
	return
}

// Login ...
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Login",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(helpers.ErrorFetchingSession())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
		return
	}

	var userAccount models.UserAccount
	err = json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Login",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(helpers.ErrorInvalidAccountPayload())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidAccountPayload()))
		return
	}

	ID, err := a.service.VerifyAccount(userAccount.Username, userAccount.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthController.Login",
			"ip":       r.RemoteAddr,
			"username": userAccount.Username,
			"err":      err.Error(),
		}).Error(helpers.ErrorInvalidUsernameOrPassword())

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidUsernameOrPassword()))
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

// Logout ...
func (a *authController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Logout",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(helpers.ErrorFetchingSession())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
		return
	}

	session.Delete()
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
