package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodePasswordReset decodes a password reset from a request and stores it in the request context.
func (s *Stack) DecodePasswordReset() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var reset models.PasswordReset
			err := json.NewDecoder(r.Body).Decode(&reset)
			if err != nil {
				s.logger.WithField("method", "middleware.DecodePasswordReset").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidPasswordResetPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextPasswordResetKey, reset)

			f(w, r.WithContext(ctx))
		}
	}
}
