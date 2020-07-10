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
	"github.com/sirupsen/logrus"
)

const (
	passwordResetCookieName = "_password_reset"
)

var (
	loginRoute         = fmt.Sprintf("http://%s/login", os.Getenv("API_HOST"))
	passwordResetRoute = fmt.Sprintf("http://%s/password/reset", os.Getenv("API_HOST"))
)

// AccountController an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_account_controller.go . AccountController
type AccountController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	ActivateAccount(w http.ResponseWriter, r *http.Request)
	RequestPasswordReset(w http.ResponseWriter, r *http.Request)
	SetPasswordResetToken(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
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

	_, err := a.service.ActivateAccount(token)
	if err != nil {
		a.logger.WithField("method", "AccountController.ActivateAccount").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorActivatingAccount()))
		return
	}

	http.Redirect(w, r, loginRoute, http.StatusSeeOther)
	return
}

func (a *accountController) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	resetRequest, ok := r.Context().Value(middleware.ContextRequestPasswordResetKey).(models.RequestPasswordReset)
	if !ok {
		a.logger.WithField("method", "AccountController.RequestPasswordReset").Error(helpers.ErrorRetrievingRequestPasswordReset())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorRequestingPasswordReset()))
		return
	}

	email, err := a.service.RequestPasswordReset(resetRequest.Username)
	if err != nil {
		if errors.Is(err, helpers.ErrorAccountNotFound()) {
			a.logger.WithFields(logrus.Fields{
				"method":   "AccountController.RequestPasswordReset",
				"username": resetRequest.Username,
			}).Info(helpers.ErrorAccountNotFound())

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.NewAPIMessage(helpers.MessageAccountWithUsernameNotFound(resetRequest.Username)))
			return
		}

		a.logger.WithFields(logrus.Fields{
			"method":   "AccountController.RequestPasswordReset",
			"username": resetRequest.Username,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorRequestingPasswordReset()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.NewAPIMessage(helpers.MessageResetPasswordEmailSent(email)))
	return
}

func (a *accountController) SetPasswordResetToken(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	t := params.Get("t")
	if t == "" || len(t) != services.PasswordResetTokenLength {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountController.SetPasswordResetToken",
			"token":  t,
		}).Error(helpers.ErrorRetrievingResetToken())

		http.Redirect(w, r, loginRoute, http.StatusSeeOther)
		return
	}

	// TODO: Ensure token is not expired
	token, err := a.service.GetPasswordResetToken(t)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountController.SetPasswordResetToken",
			"token":  t,
		}).Error(err.Error())

		http.Redirect(w, r, loginRoute, http.StatusSeeOther)
		return
	}

	cookie := http.Cookie{
		Name:     passwordResetCookieName,
		Value:    token.ResetToken,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, passwordResetRoute, http.StatusSeeOther)
	return
}

func (a *accountController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(passwordResetCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			a.logger.WithField("method", "AccountController.ResetPassword").Info(err.Error())

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorRetrievingPasswordReset()))
			return
		}

		a.logger.WithField("method", "AccountController.ResetPassword").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorRetrievingPasswordReset()))
		return
	}

	reset, ok := r.Context().Value(middleware.ContextPasswordResetKey).(models.PasswordReset)
	if !ok {
		a.logger.WithFields(logrus.Fields{
			"method": "AccountController.ResetPassword",
			"token":  cookie.Value,
		}).Error(helpers.ErrorRetrievingPasswordReset())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorRetrievingPasswordReset()))
		return
	}

	_, err = a.service.ResetPassword(cookie.Value, reset.Password)
	if err != nil {
		if errors.Is(err, helpers.ErrorResetTokenExpiredOrInvalid()) {
			a.logger.WithFields(logrus.Fields{
				"method": "AccountController.ResetPassword",
				"token":  cookie.Value,
			}).Info(err.Error())

			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorResetTokenExpiredOrInvalid()))
			return
		}

		a.logger.WithFields(logrus.Fields{
			"method": "AccountController.ResetPassword",
			"token":  cookie.Value,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorResettingPassword()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.NewAPIMessage(helpers.MessagePasswordUpdated()))
	return
}
