package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/sirupsen/logrus"
)

func (m *MiddlewareStack) DecodeExpense() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m.logger.WithFields(logrus.Fields{
				"method": "middleware.DecodeExpense",
				"ip":     r.RemoteAddr,
			})

			var expense models.Expense
			err := json.NewDecoder(r.Body).Decode(&expense)
			if err != nil {
				m.logger.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Error(controllers.ErrInvalidExpensePayload)

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.NewAPIError(controllers.ErrInvalidExpensePayload))
				return
			}

			ctx := context.WithValue(r.Context(), "expense", expense)

			f(w, r.WithContext(ctx))
		}
	}
}
