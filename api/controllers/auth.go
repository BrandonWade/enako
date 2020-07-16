package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/csrf"
	"github.com/sirupsen/logrus"
)

// AuthController an interface for wotking with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_auth_controller.go . AuthController
type AuthController interface {
	CSRF(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.AccountService
}

// NewAuthController returns a new instance of an AuthController.
func NewAuthController(logger *logrus.Logger, store helpers.CookieStorer, service services.AccountService) AuthController {
	return &authController{
		logger,
		store,
		service,
	}
}

// CSRF returns a new anti-CSRF token for the SPA frontend.
func (a *authController) CSRF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.WriteHeader(http.StatusOK)
	return
}

// Login creates a new session for an account.
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingSession()))
		return
	}

	var account models.Account
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidAccountPayload()))
		return
	}

	ID, err := a.service.VerifyAccount(account.Email, account.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AuthController.Login",
			"email":  account.Email,
		}).Error(err.Error())

		w.WriteHeader(http.StatusUnauthorized)
		if errors.Is(err, helpers.ErrorActivationEmailResent()) {
			json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessageActivationEmailSent(account.Email)))
		} else if errors.Is(err, helpers.ErrorAccountNotActivated()) {
			json.NewEncoder(w).Encode(models.MessagesFromErrors(err))
		} else {
			json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidEmailOrPassword()))
		}

		return
	}

	session.Set("authenticated", true)
	session.Set("account_id", ID)
	session.Save(r, w)

	account.ID = ID
	account.Password = ""

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
	return
}

// Logout deletes the current account session.
func (a *authController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithField("method", "AuthController.Logout").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingSession()))
		return
	}

	session.Delete()
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
