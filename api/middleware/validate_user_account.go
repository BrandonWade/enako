package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
)

// ValidateUserAccount ...
func (m *MiddlewareStack) ValidateUserAccount() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			userAccount := r.Context().Value(ContextUserAccountKey).(models.UserAccount)

			if err := validator.Validate(userAccount); err != nil {
				m.logger.WithFields(logrus.Fields{
					"method": "middleware.ValidateUserAccount",
					"ip":     r.RemoteAddr,
					"err":    err.Error(),
				}).Error(err)

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(err))
				return
			}

			if userAccount.Password != userAccount.ConfirmPassword {
				m.logger.WithFields(logrus.Fields{
					"method": "AuthController.CreateAccount",
					"ip":     r.RemoteAddr,
				}).Error(helpers.ErrorPasswordsDoNotMatch())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorPasswordsDoNotMatch()))
				return
			}

			f(w, r)
		}
	}
}
