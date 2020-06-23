package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/csrf"
	"github.com/sirupsen/logrus"
)

// AuthController an interface for wotking with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_auth_controller.go . AuthController
type AuthController interface {
	CSRF(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	ActivateAccount(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.AuthService
}

// NewAuthController returns a new instance of an AuthController.
func NewAuthController(logger *logrus.Logger, store helpers.CookieStorer, service services.AuthService) AuthController {
	return &authController{
		logger,
		store,
		service,
	}
}

// CSRF returns a new anti-CSRF token for our SPA frontend.
func (a *authController) CSRF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.WriteHeader(http.StatusOK)
	return
}

// CreateAccount creates a new account.
func (a *authController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	createAccount, ok := r.Context().Value(middleware.ContextCreateAccountKey).(models.CreateAccount)
	if !ok {
		a.logger.WithField("method", "AuthController.CreateAccount").Error(helpers.ErrorRetrievingAccount())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	id, token, err := a.service.CreateAccount(createAccount.Username, createAccount.Email, createAccount.Password)
	if err != nil {
		a.logger.WithField("method", "AuthController.CreateAccount").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	createAccount.ID = id
	createAccount.Password = ""
	createAccount.ConfirmPassword = ""
	createAccount.ActivationLink = fmt.Sprintf("%s/api/v1/accounts/activate?t=%s", os.Getenv("API_HOST"), token) // TODO: Remove this; the link should be sent to the provided email

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createAccount)
	return
}

// ActivateAccount activates a newly created account.
func (a *authController) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("t")
	if len(token) != services.ActivationTokenLength {
		a.logger.WithField("method", "AuthController.ActivateAccount").Error(helpers.ErrorInvalidActivationToken())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorActivatingAccount()))
		return
	}

	success, err := a.service.ActivateAccount(token)
	if !success {
		if err != nil {
			a.logger.WithField("method", "AuthController.ActivateAccount").Error(err.Error())
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorActivatingAccount()))
		return
	}

	login := fmt.Sprintf("http://%s/login", os.Getenv("API_HOST"))
	http.Redirect(w, r, login, http.StatusSeeOther)
	return
}

// Login creates a new session for an account.
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
		return
	}

	var account models.Account
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidAccountPayload()))
		return
	}

	ID, err := a.service.VerifyAccount(account.Username, account.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthController.Login",
			"username": account.Username,
		}).Error(err.Error())

		errMsg := helpers.ErrorInvalidUsernameOrPassword()
		if errors.Is(err, helpers.ErrorAccountNotActivated()) {
			errMsg = helpers.ErrorAccountNotActivated()
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.NewAPIError(errMsg))
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
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
		return
	}

	session.Delete()
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
