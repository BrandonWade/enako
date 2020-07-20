package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
)

const (
	// PasswordResetCookieName the name for the cookie used to hold a password reset token
	PasswordResetCookieName = "_password_reset"

	// PasswordResetCookieMaxAge the time to live in seconds for the password reset token cookie
	PasswordResetCookieMaxAge = 86400 // 24 hours
)

// PasswordResetController an interface for working with password resets.
//go:generate counterfeiter -o fakes/fake_password_reset_controller.go . PasswordResetController
type PasswordResetController interface {
	RequestPasswordReset(w http.ResponseWriter, r *http.Request)
	SetPasswordResetToken(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

type passwordResetController struct {
	logger  *logrus.Logger
	service services.PasswordResetService
}

// NewPasswordResetController returns a new instance of an PasswordResetController.
func NewPasswordResetController(logger *logrus.Logger, service services.PasswordResetService) PasswordResetController {
	return &passwordResetController{
		logger,
		service,
	}
}

func (a *passwordResetController) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	resetRequest, ok := r.Context().Value(middleware.ContextRequestPasswordResetKey).(models.RequestPasswordReset)
	if !ok {
		a.logger.WithField("method", "PasswordResetController.RequestPasswordReset").Error(helpers.ErrorRetrievingRequestPasswordReset())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRequestingPasswordReset()))
		return
	}

	email, err := a.service.RequestPasswordReset(resetRequest.Email)
	if err != nil {
		if errors.Is(err, helpers.ErrorAccountNotFound()) {
			a.logger.WithFields(logrus.Fields{
				"method": "PasswordResetController.RequestPasswordReset",
				"email":  resetRequest.Email,
			}).Info(helpers.ErrorAccountNotFound())

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessageAccountWithEmailNotFound(resetRequest.Email)))
			return
		}

		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetController.RequestPasswordReset",
			"email":  resetRequest.Email,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRequestingPasswordReset()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessageResetPasswordEmailSent(email)))
	return
}

func (a *passwordResetController) SetPasswordResetToken(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	t := params.Get("t")
	if t == "" || len(t) != services.PasswordResetTokenLength {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetController.SetPasswordResetToken",
			"token":  t,
		}).Error(helpers.ErrorRetrievingResetToken())

		http.Redirect(w, r, helpers.LoginRoute, http.StatusSeeOther)
		return
	}

	err := a.service.VerifyPasswordResetToken(t)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetController.SetPasswordResetToken",
			"token":  t,
		}).Error(err.Error())

		// TODO: Show message indicating reset token was invalid

		http.Redirect(w, r, helpers.ForgotPasswordRoute, http.StatusSeeOther)
		return
	}

	cookie := http.Cookie{
		Name:     PasswordResetCookieName,
		Value:    t,
		Path:     "/",
		MaxAge:   PasswordResetCookieMaxAge,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, helpers.PasswordResetRoute, http.StatusSeeOther)
	return
}

func (a *passwordResetController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(PasswordResetCookieName)
	if err != nil {
		a.logger.WithField("method", "PasswordResetController.ResetPassword").Info(err.Error())

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRetrievingPasswordReset()))
		return
	}

	reset, ok := r.Context().Value(middleware.ContextPasswordResetKey).(models.PasswordReset)
	if !ok {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetController.ResetPassword",
			"token":  cookie.Value,
		}).Error(helpers.ErrorRetrievingPasswordReset())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRetrievingPasswordReset()))
		return
	}

	_, err = a.service.ResetPassword(cookie.Value, reset.Password)
	if err != nil {
		if errors.Is(err, helpers.ErrorResetTokenExpiredOrInvalid()) {
			a.logger.WithFields(logrus.Fields{
				"method": "PasswordResetController.ResetPassword",
				"token":  cookie.Value,
			}).Info(err.Error())

			// TODO: Show message indicating reset token was invalid

			http.Redirect(w, r, helpers.ForgotPasswordRoute, http.StatusSeeOther)
			return
		}

		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetController.ResetPassword",
			"token":  cookie.Value,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorResettingPassword()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessagePasswordUpdated()))
	return
}
