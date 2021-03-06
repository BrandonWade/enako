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

// AccountController an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_account_controller.go . AccountController
type AccountController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	ActivateAccount(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	RequestChangeEmail(w http.ResponseWriter, r *http.Request)
}

type accountController struct {
	logger  *logrus.Logger
	service services.AccountService
}

// NewAccountController returns a new instance of an AccountController.
func NewAccountController(logger *logrus.Logger, service services.AccountService) AccountController {
	return &accountController{
		logger,
		service,
	}
}

// RegisterUser creates a new account.
func (a *accountController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	createAccount, ok := r.Context().Value(middleware.ContextCreateAccountKey).(models.CreateAccount)
	if !ok {
		a.logger.WithField("method", "AccountController.RegisterUser").Error(helpers.ErrorRetrievingAccount())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorCreatingAccount()))
		return
	}

	_, err := a.service.RegisterUser(createAccount.Email, createAccount.Password)
	if err != nil {
		a.logger.WithField("method", "AccountController.RegisterUser").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorCreatingAccount()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessageActivationEmailSent(createAccount.Email)))
	return
}

// ActivateAccount activates a newly created account.
func (a *accountController) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("t")
	if len(token) != services.ActivationTokenLength {
		a.logger.WithField("method", "AccountController.ActivateAccount").Error(helpers.ErrorInvalidActivationToken())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorActivatingAccount()))
		return
	}

	_, err := a.service.ActivateAccount(token)
	if err != nil {
		a.logger.WithField("method", "AccountController.ActivateAccount").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorActivatingAccount()))
		return
	}

	http.Redirect(w, r, helpers.LoginRoute, http.StatusSeeOther)
	return
}

// ChangePassword changes the password for the account in the current session.
func (a *accountController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		a.logger.WithField("method", "AccountController.ChangePassword").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorChangingPassword()))
		return
	}

	changePassword, ok := r.Context().Value(middleware.ContextChangePasswordKey).(models.ChangePassword)
	if !ok {
		a.logger.WithField("method", "AccountController.ChangePassword").Error(helpers.ErrorRetrievingChangePassword())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorChangingPassword()))
		return
	}

	_, err := a.service.ChangePassword(accountID, changePassword.CurrentPassword, changePassword.NewPassword)
	if err != nil {
		a.logger.WithField("method", "AccountController.ChangePassword").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorChangingPassword()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessagePasswordUpdated()))
	return
}

//RequestChangeEmail initiates a request to change the email for the account in the current session.
func (a *accountController) RequestChangeEmail(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement

	w.WriteHeader(http.StatusOK)
	return
}
