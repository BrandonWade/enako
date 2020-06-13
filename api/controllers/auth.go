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

// AuthController an interface for wotking with user accounts and sessions.
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

// NewAuthController returns a new instance of an AuthController.
func NewAuthController(logger *logrus.Logger, store helpers.CookieStorer, service services.AuthService) AuthController {
	return &authController{
		logger,
		store,
		service,
	}
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

	ID, err := a.service.CreateAccount(createAccount.Username, createAccount.Email, createAccount.Password)
	if err != nil {
		a.logger.WithField("method", "AuthController.CreateAccount").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	createAccount.ID = ID
	createAccount.Password = ""
	createAccount.ConfirmPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createAccount)
	return
}

// Login creates a new session for a user account.
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, helpers.SessionCookieName)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
		return
	}

	var userAccount models.UserAccount
	err = json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		a.logger.WithField("method", "AuthController.Login").Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidAccountPayload()))
		return
	}

	ID, err := a.service.VerifyAccount(userAccount.Username, userAccount.Password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "AuthController.Login",
			"username": userAccount.Username,
		}).Error(err.Error())

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

// Logout deletes the current user account session.
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
