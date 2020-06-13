package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidateUserAccount checks whether a decoded user account in a request is valid.
func (m *MiddlewareStack) ValidateUserAccount() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			userAccount, ok := r.Context().Value(ContextUserAccountKey).(models.UserAccount)
			if !ok {
				m.logger.WithField("method", "middleware.ValidateUserAccount").Error(helpers.ErrorRetrievingAccount())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidAccountPayload()))
				return
			}

			if err := validator.Validate(userAccount); err != nil {
				m.logger.WithField("method", "middleware.ValidateUserAccount").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(err))
				return
			}

			if userAccount.Password != userAccount.ConfirmPassword {
				m.logger.WithField("method", "middleware.ValidateUserAccount").Info(helpers.ErrorPasswordsDoNotMatch())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorPasswordsDoNotMatch()))
				return
			}

			f(w, r)
		}
	}
}
