package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeCreateAccount decodes a create account payload from a request and stores it in the request context.
func (s *Stack) DecodeCreateAccount() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var createAccount models.CreateAccount
			err := json.NewDecoder(r.Body).Decode(&createAccount)
			if err != nil {
				s.logger.WithField("method", "middleware.DecodeCreateAccount").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidAccountPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextCreateAccountKey, createAccount)

			f(w, r.WithContext(ctx))
		}
	}
}
