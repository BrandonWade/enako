package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeUserAccount ...
func (m *MiddlewareStack) DecodeUserAccount() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var userAccount models.UserAccount
			err := json.NewDecoder(r.Body).Decode(&userAccount)
			if err != nil {
				m.logger.WithField("method", "middleware.DecodeUserAccount").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidAccountPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextUserAccountKey, userAccount)

			f(w, r.WithContext(ctx))
		}
	}
}
