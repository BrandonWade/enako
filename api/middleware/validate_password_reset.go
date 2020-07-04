package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidatePasswordReset checks whether a decoded request password reset in the payload is valid.
func (m *MiddlewareStack) ValidatePasswordReset() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			reset, ok := r.Context().Value(ContextPasswordResetKey).(models.PasswordReset)
			if !ok {
				m.logger.WithField("method", "middleware.ValidatePasswordReset").Error(helpers.ErrorRetrievingPasswordReset())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidPasswordResetPayload()))
				return
			}

			if err := validator.Validate(reset); err != nil {
				m.logger.WithField("method", "middleware.ValidatePasswordReset").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(err))
				return
			}

			f(w, r)
		}
	}
}
