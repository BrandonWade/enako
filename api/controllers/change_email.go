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

// ChangeEmailController an interface for changing the email associated with an account.
//go:generate counterfeiter -o fakes/fake_change_email_controller.go . ChangeEmailController
type ChangeEmailController interface {
	RequestEmailChange(w http.ResponseWriter, r *http.Request)
}

type changeEmailController struct {
	logger  *logrus.Logger
	service services.ChangeEmailService
}

// NewChangeEmailController returns a new instance of an ChangeEmailController.
func NewChangeEmailController(logger *logrus.Logger, service services.ChangeEmailService) ChangeEmailController {
	return &changeEmailController{
		logger,
		service,
	}
}

// RequestEmailChange initiates a request to change the email for the account in the current session.
func (a *changeEmailController) RequestEmailChange(w http.ResponseWriter, r *http.Request) {
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		a.logger.WithField("method", "ChangeEmailController.RequestEmailChange").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRequestingEmailChange()))
		return
	}

	email, err := a.service.RequestEmailChange(accountID)
	if err != nil {
		a.logger.WithField("method", "ChangeEmailController.RequestEmailChange").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorRequestingEmailChange()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.MessagesFromStrings(helpers.MessageChangeEmailEmailSent(email)))
	return
}
