package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
)

// DecodeExpense decodes an expense from a request and stores it in the request context.
func (m *MiddlewareStack) DecodeExpense() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var expense models.Expense
			err := json.NewDecoder(r.Body).Decode(&expense)
			if err != nil {
				m.logger.WithField("method", "middleware.DecodeExpense").Info(err.Error())

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpensePayload()))
				return
			}

			ctx := context.WithValue(r.Context(), ContextExpenseKey, expense)

			f(w, r.WithContext(ctx))
		}
	}
}
