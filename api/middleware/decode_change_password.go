package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeChangePassword decodes a change password payload from a request and stores it in the request context.
func (s *Stack) DecodeChangePassword() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var changePassword models.ChangePassword
			err := json.NewDecoder(r.Body).Decode(&changePassword)
			if err != nil {
				s.logger.WithField("method", "middleware.DecodeChangePassword").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidChangePasswordPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextChangePasswordKey, changePassword)

			f(w, r.WithContext(ctx))
		}
	}
}
