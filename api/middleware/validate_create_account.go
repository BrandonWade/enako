package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidateCreateAccount checks whether a decoded CreateAccount payload in a request is valid.
func (m *MiddlewareStack) ValidateCreateAccount() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			createAccount, ok := r.Context().Value(ContextCreateAccountKey).(models.CreateAccount)
			if !ok {
				m.logger.WithField("method", "middleware.ValidateCreateAccount").Error(helpers.ErrorRetrievingAccount())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidAccountPayload()))
				return
			}

			if err := validator.Validate(createAccount); err != nil {
				m.logger.WithField("method", "middleware.ValidateCreateAccount").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(err))
				return
			}

			if createAccount.Password != createAccount.ConfirmPassword {
				m.logger.WithField("method", "middleware.ValidateCreateAccount").Info(helpers.ErrorPasswordsDoNotMatch())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorPasswordsDoNotMatch()))
				return
			}

			f(w, r)
		}
	}
}
