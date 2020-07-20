package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeLogin decodes the login information from a request into an account model and stores it in the request context.
func (s *Stack) DecodeLogin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var account models.Account
			err := json.NewDecoder(r.Body).Decode(&account)
			if err != nil {
				s.logger.WithField("method", "middleware.DecodeLogin").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidAccountPayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextLoginKey, account)

			f(w, r.WithContext(ctx))
		}
	}
}
