package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	validator "gopkg.in/validator.v2"
)

var (
	ErrFetchingExpenses      = errors.New("error fetching expense list")
	ErrInvalidExpensePayload = errors.New("invalid expense payload")
	ErrCreatingExpense       = errors.New("error creating expense")
	ErrInvalidExpenseID      = errors.New("invalid expense id")
	ErrUpdatingExpense       = errors.New("error updating expense")
	ErrNoExpensesUpdated     = errors.New("no expenses were updated")
	ErrDeletingExpense       = errors.New("error deleting expense")
	ErrNoExpensesDeleted     = errors.New("no expenses were deleted")
)

// ExpenseController the interface for expense related APIs
//go:generate counterfeiter -o fakes/fake_expense_controller.go . ExpenseController
type ExpenseController interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, e *models.Expense)
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
		}).Error(ErrFetchingExpenses)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrFetchingExpenses))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
	return
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, expense *models.Expense) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	if err := validator.Validate(expense); err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "ExpenseController.CreateExpense",
			"err":    err.Error(),
		}).Error(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(err))
		return
	}

	ID, err := e.service.CreateExpense(userAccountID, expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error(ErrCreatingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrCreatingExpense))
		return
	}

	expense.ID = ID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
	return
}

func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         params["id"],
			"err":        err.Error(),
		}).Error(ErrInvalidExpenseID)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidExpenseID))
		return
	}

	expense := models.Expense{}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(ErrInvalidExpensePayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidExpensePayload))
		return
	}

	if err = validator.Validate(expense); err != nil {
		e.logger.WithFields(logrus.Fields{
			"method": "ExpenseController.UpdateExpense",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(err))
		return
	}

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(ErrUpdatingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrUpdatingExpense))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Warn(ErrNoExpensesUpdated)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrNoExpensesUpdated))
		return
	}

	expense.ID = ID

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
		}).Error(ErrInvalidExpenseID)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrInvalidExpenseID))
		return
	}

	count, err := e.service.DeleteExpense(ID, userAccountID)
	if err != nil {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(ErrDeletingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrDeletingExpense))
		return
	}

	if count == 0 {
		e.logger.WithFields(logrus.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
		}).Warn(ErrNoExpensesDeleted)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrNoExpensesDeleted))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
