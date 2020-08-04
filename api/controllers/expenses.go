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
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.GetExpenses").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpensePayload()))
		return
	}

	expenses, err := e.service.GetExpenses(accountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.GetExpenses",
			"account ID": accountID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingExpenses()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
	return
}

// CreateExpense creates a new expense.
func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.CreateExpenses").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpensePayload()))
		return
	}

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": accountID,
		}).Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorCreatingExpense()))
		return
	}

	ID, err := e.service.CreateExpense(accountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": accountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorCreatingExpense()))
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
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.UpdateExpenses").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorUpdatingExpense()))
		return
	}

	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": accountID,
			"id":         params["id"],
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpenseID()))
		return
	}

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": accountID,
			"id":         ID,
		}).Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorUpdatingExpense()))
		return
	}

	count, err := e.service.UpdateExpense(ID, accountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": accountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorUpdatingExpense()))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": accountID,
			"id":         ID,
		}).Warn(helpers.ErrorNoExpensesUpdated())

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorNoExpensesUpdated()))
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
	accountID, ok := r.Context().Value(middleware.ContextAccountIDKey).(int64)
	if !ok {
		e.logger.WithField("method", "ExpenseController.DeleteExpenses").Error(helpers.ErrorRetrievingAccountID())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpensePayload()))
		return
	}

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": accountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorInvalidExpenseID()))
		return
	}

	count, err := e.service.DeleteExpense(ID, accountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": accountID,
			"id":         ID,
		}).Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorDeletingExpense()))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": accountID,
			"id":         ID,
		}).Warn(helpers.ErrorNoExpensesDeleted())

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorNoExpensesDeleted()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
