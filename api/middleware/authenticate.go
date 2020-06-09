package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"

	"github.com/BrandonWade/enako/api/models"
	"github.com/sirupsen/logrus"
)

// Authenticate ...
func (m *MiddlewareStack) Authenticate() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m.logger.WithFields(logrus.Fields{
				"method": "middleware.Authenticate",
				"ip":     r.RemoteAddr,
			})

			authenticated, err := m.store.IsAuthenticated(r)
			if err != nil {
				m.logger.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Error(helpers.ErrorFetchingSession())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingSession()))
				return
			}

			if !authenticated {
				m.logger.Info(helpers.ErrorUserNotAuthenticated())

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUserNotAuthenticated()))
				return
			}

			f(w, r)
		}
	}
}
