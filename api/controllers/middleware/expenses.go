package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/sirupsen/logrus"
)

// DecodeExpense ...
func DecodeExpense(logger *logrus.Logger, next func(w http.ResponseWriter, e *models.Expense)) http.HandlerFunc {
	logger.WithFields(logrus.Fields{"function": "middleware.DecodeExpense"})

	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		err := json.NewDecoder(r.Body).Decode(&expense)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error(controllers.ErrInvalidExpensePayload)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.NewAPIError(controllers.ErrInvalidExpensePayload))
			return
		}

		next(w, &expense)
	}
}
