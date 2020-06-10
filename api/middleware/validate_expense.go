package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
)

// ValidateExpense ...
func (m *MiddlewareStack) ValidateExpense() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			expense, ok := r.Context().Value(ContextExpenseKey).(models.Expense)
			if !ok {
				m.logger.Error(helpers.ErrorRetrievingExpense())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingExpense()))
				return
			}

			if err := validator.Validate(expense); err != nil {
				m.logger.WithFields(logrus.Fields{
					"method": "middleware.ValidateExpense",
					"err":    err.Error(),
				}).Error(err)

				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(models.NewAPIError(err))
				return
			}

			f(w, r)
		}
	}
}
