package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"gopkg.in/validator.v2"
)

// ValidateExpense checks whether a decoded expense in a request is valid.
func (m *MiddlewareStack) ValidateExpense() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			expense, ok := r.Context().Value(ContextExpenseKey).(models.Expense)
			if !ok {
				m.logger.WithField("method", "middleware.ValidateExpense").Error(helpers.ErrorRetrievingExpense())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpensePayload()))
				return
			}

			if err := validator.Validate(expense); err != nil {
				m.logger.WithField("method", "middleware.ValidateExpense").Info(err.Error())

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.MessagesFromErrors(err))
				return
			}

			f(w, r)
		}
	}
}
