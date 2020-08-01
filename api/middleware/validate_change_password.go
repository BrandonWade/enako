package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidateChangePassword checks whether a decoded ChangePassword payload in a request is valid.
func (s *Stack) ValidateChangePassword() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			changePassword, ok := r.Context().Value(ContextChangePasswordKey).(models.ChangePassword)
			if !ok {
				s.logger.WithField("method", "middleware.ValidateChangePassword").Error(helpers.ErrorRetrievingChangePassword())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidChangePasswordPayload()))
				return
			}

			if err := validator.Validate(changePassword); err != nil {
				s.logger.WithField("method", "middleware.ValidateChangePassword").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(err))
				return
			}

			if changePassword.CurrentPassword == changePassword.NewPassword {
				s.logger.WithField("method", "middleware.ValidateChangePassword").Info(helpers.ErrorPasswordsShouldNotMatch())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorPasswordsShouldNotMatch()))
				return
			}

			if changePassword.NewPassword != changePassword.ConfirmPassword {
				s.logger.WithField("method", "middleware.ValidateChangePassword").Info(helpers.ErrorPasswordsDoNotMatch())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorPasswordsDoNotMatch()))
				return
			}

			f(w, r)
		}
	}
}
