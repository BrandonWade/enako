package controllers

import (
	"net/http"

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
	// TODO: Implement

	w.WriteHeader(http.StatusOK)
	return
}
