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

// ExpenseController the interface for expense related APIs
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

// NewExpenseController the constructor for a new ExpenseController
func NewExpenseController(logger *logrus.Logger, store helpers.CookieStorer, service services.ExpenseService) ExpenseController {
	return &expenseController{
		logger,
		store,
		service,
	}
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	expenses, err := e.service.GetExpenses(userAccountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.GetExpenses",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error(helpers.ErrorFetchingExpenses())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingExpenses()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
	return
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	e.logger.WithFields(logrus.Fields{
		"method":     "ExpenseController.CreateExpense",
		"account ID": userAccountID,
	})

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorCreatingExpense()))
		return
	}

	ID, err := e.service.CreateExpense(userAccountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(helpers.ErrorCreatingExpense())

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

func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	e.logger.WithFields(logrus.Fields{
		"method":     "ExpenseController.UpdateExpense",
		"account ID": userAccountID,
	})

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"id":  params["id"],
			"err": err.Error(),
		}).Error(helpers.ErrorInvalidExpenseID())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorInvalidExpenseID()))
		return
	}

	expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)
	if !ok {
		e.logger.Error(helpers.ErrorRetrievingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUpdatingExpense()))
		return
	}

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"id":  ID,
			"err": err.Error(),
		}).Error(helpers.ErrorUpdatingExpense())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorUpdatingExpense()))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"id": ID,
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

func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(helpers.ErrorInvalidExpenseID())

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
			"err":        err.Error(),
		}).Error(helpers.ErrorDeletingExpense())

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
