package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ExpenseController the interface for working with expenses.
//go:generate counterfeiter -o fakes/fake_expense_controller.go . ExpenseController
type ExpenseController interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
}

type expenseController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.ExpenseService
}

// NewExpenseController returns a new instance of an ExpenseController.
func NewExpenseController(logger *logrus.Logger, store helpers.CookieStorer, service services.ExpenseService) ExpenseController {
	return &expenseController{
		logger,
		store,
		service,
	}
}

// GetExpenses returns the list of expenses.
func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userAccountID, ok := r.Context().Value(middleware.ContextUserAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.GetExpenses").Error(helpers.ErrorRetrievingUserAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpensePayload()))
		return
	}

	expenses, err := e.service.GetExpenses(userAccountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.GetExpenses",
			"account ID": userAccountID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingExpenses()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
	return
}

// CreateExpense creates a new expense.
func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID, ok := r.Context().Value(middleware.ContextUserAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.CreateExpenses").Error(helpers.ErrorRetrievingUserAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpensePayload()))
		return
	}

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
		}).Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingExpense()))
		return
	}

	ID, err := e.service.CreateExpense(userAccountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingExpense()))
		return
	}

	expense.ID = ID
	expense.Amount = expense.Amount / 100

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
	return
}

// UpdateExpense updates the expense with the given id.
func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID, ok := r.Context().Value(middleware.ContextUserAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.UpdateExpenses").Error(helpers.ErrorRetrievingUserAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpensePayload()))
		return
	}

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         params["id"],
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpenseID()))
		return
	}

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUpdatingExpense()))
		return
	}

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUpdatingExpense()))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Warn(helpers.ErrorNoExpensesUpdated())

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorNoExpensesUpdated()))
		return
	}

	expense.ID = ID
	expense.Amount = expense.Amount / 100

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
	return
}

// DeletesExpense deletes the expense with the submitted id.
func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID, ok := r.Context().Value(middleware.ContextUserAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.DeleteExpenses").Error(helpers.ErrorRetrievingUserAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpensePayload()))
		return
	}

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpenseID()))
		return
	}

	count, err := e.service.DeleteExpense(ID, userAccountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorDeletingExpense()))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Warn(helpers.ErrorNoExpensesDeleted())

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorNoExpensesDeleted()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
