package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"

	"github.com/BrandonWade/enako/api/models"
)

// Authenticate checks whether a valid session exists for the request.
func (m *MiddlewareStack) Authenticate() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authenticated, err := m.store.IsAuthenticated(r)
			if err != nil {
				m.logger.WithField("method", "middleware.Authenticate").Error(err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingSession()))
				return
			}

			if !authenticated {
				m.logger.WithField("method", "middleware.Authenticate").Info(helpers.ErrorUserNotAuthenticated())

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorUserNotAuthenticated()))
				return
			}

			session, err := m.store.Get(r, helpers.SessionCookieName)
			if err != nil {
				m.logger.WithField("method", "AuthController.Login").Error(err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingSession()))
				return
			}

			accountID := session.Get("account_id")
			ctx := context.WithValue(r.Context(), ContextAccountIDKey, accountID)

			f(w, r.WithContext(ctx))
		}
	}
}
