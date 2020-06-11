package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"

	"github.com/BrandonWade/enako/api/models"
)

// Authenticate ...
func (m *MiddlewareStack) Authenticate() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authenticated, err := m.store.IsAuthenticated(r)
			if err != nil {
				m.logger.WithField("method", "middleware.Authenticate").Error(err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
				return
			}

			if !authenticated {
				m.logger.WithField("method", "middleware.Authenticate").Info(helpers.ErrorUserNotAuthenticated())

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUserNotAuthenticated()))
				return
			}

			// TODO: Inject user account ID from cookie into context

			f(w, r)
		}
	}
}
