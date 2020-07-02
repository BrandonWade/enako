package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidateRequestPasswordReset checks whether a decoded ResetPasswordRequest payload in a request is valid.
func (m *MiddlewareStack) ValidateRequestPasswordReset() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			requestPasswordReset, ok := r.Context().Value(ContextRequestPasswordResetKey).(models.RequestPasswordReset)
			if !ok {
				m.logger.WithField("method", "middleware.ValidateRequestPasswordReset").Error(helpers.ErrorRetrievingRequestPasswordReset())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidRequestPasswordResetPayload()))
				return
			}

			if err := validator.Validate(requestPasswordReset); err != nil {
				m.logger.WithField("method", "middleware.ValidateRequestPasswordReset").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(err))
				return
			}

			f(w, r)
		}
	}
}
