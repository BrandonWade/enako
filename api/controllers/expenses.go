package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
	validator "gopkg.in/validator.v2"
)

var (
	errFetchingExpenses      = errors.New("error fetching expense list")
	errInvalidExpensePayload = errors.New("invalid expense payload")
	errCreatingExpense       = errors.New("error creating expense")
	errInvalidExpenseID      = errors.New("invalid expense id")
	errUpdatingExpense       = errors.New("error updating expense")
	errNoExpensesUpdated     = errors.New("no expenses were updated")
	errDeletingExpense       = errors.New("error deleting expense")
	errNoExpensesDeleted     = errors.New("no expenses were deleted")
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
	service services.ExpenseService
}

// NewExpenseController the constructor for a new ExpenseController
func NewExpenseController(service services.ExpenseService) ExpenseController {
	return &expenseController{
		service,
	}
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	expenses, err := e.service.GetExpenses(userAccountID)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.GetExpenses",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error(errFetchingExpenses)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errFetchingExpenses))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
	return
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	var expense models.UserExpense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error(errInvalidExpensePayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(errInvalidExpensePayload))
		return
	}

	if err = validator.Validate(expense); err != nil {
		log.WithFields(log.Fields{
			"method": "ExpenseController.CreateExpense",
			"ip":     r.RemoteAddr,
			"err":    err.Error(),
		}).Error(err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(err))
		return
	}

	ID, err := e.service.CreateExpense(userAccountID, &expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error(errCreatingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errCreatingExpense))
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
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         params["id"],
			"err":        err.Error(),
		}).Error(errInvalidExpenseID)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(errInvalidExpenseID))
		return
	}

	expense := models.UserExpense{}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(errInvalidExpensePayload)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(errInvalidExpensePayload))
		return
	}

	// TODO: Validate inputs

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(errUpdatingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errUpdatingExpense))
		return
	}

	if count == 0 {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Warn(errNoExpensesUpdated)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.NewAPIError(errNoExpensesUpdated))
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
		log.WithFields(log.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(errInvalidExpenseID)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(errInvalidExpenseID))
		return
	}

	count, err := e.service.DeleteExpense(ID, userAccountID)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error(errDeletingExpense)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errDeletingExpense))
		return
	}

	if count == 0 {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Warn(errNoExpensesDeleted)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.NewAPIError(errNoExpensesDeleted))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
