package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeRequestPasswordReset decodes a requested password reset from a request and stores it in the request context.
func (s *Stack) DecodeRequestPasswordReset() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var resetRequest models.RequestPasswordReset
			err := json.NewDecoder(r.Body).Decode(&resetRequest)
			if err != nil {
				s.logger.WithField("method", "middleware.DecodeRequestPasswordReset").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidRequestPasswordResetPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextRequestPasswordResetKey, resetRequest)

			f(w, r.WithContext(ctx))
		}
	}
}
