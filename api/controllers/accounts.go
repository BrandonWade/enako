package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
)

// AccountController an interface for wotking with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_account_controller.go . AccountController
type AccountController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	ActivateAccount(w http.ResponseWriter, r *http.Request)
}

type accountController struct {
	logger  *logrus.Logger
	service services.AuthService
}

// NewAccountController returns a new instance of an AccountController.
func NewAccountController(logger *logrus.Logger, service services.AuthService) AccountController {
	return &accountController{
		logger,
		service,
	}
}

// RegisterUser creates an account, generates an activation token, and sends an activation email.
func (a *accountController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	createAccount, ok := r.Context().Value(middleware.ContextCreateAccountKey).(models.CreateAccount)
	if !ok {
		a.logger.WithField("method", "AccountController.RegisterUser").Error(helpers.ErrorRetrievingAccount())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	id, err := a.service.RegisterUser(createAccount.Username, createAccount.Email, createAccount.Password)
	if err != nil {
		a.logger.WithField("method", "AccountController.RegisterUser").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingAccount()))
		return
	}

	createAccount.ID = id
	createAccount.Password = ""
	createAccount.ConfirmPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createAccount)
	return
}

// ActivateAccount activates a newly created account.
func (a *accountController) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("t")
	if len(token) != services.ActivationTokenLength {
		a.logger.WithField("method", "AccountController.ActivateAccount").Error(helpers.ErrorInvalidActivationToken())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorActivatingAccount()))
		return
	}

	success, err := a.service.ActivateAccount(token)
	if !success {
		if err != nil {
			a.logger.WithField("method", "AccountController.ActivateAccount").Error(err.Error())
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorActivatingAccount()))
		return
	}

	login := fmt.Sprintf("http://%s/login", os.Getenv("API_HOST"))
	http.Redirect(w, r, login, http.StatusSeeOther)
	return
}
