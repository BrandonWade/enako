package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
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
		}).Error("error fetching expense list")

		// TODO: Return an error response
	}

	json.NewEncoder(w).Encode(expenses)
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
		}).Error("error decoding expense payload")

		// TODO: Return an error response
	}

	ID, err := e.service.CreateExpense(userAccountID, &expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.CreateExpense",
			"account ID": userAccountID,
			"err":        err.Error(),
		}).Error("error creating expense")

		// TODO: Return an error response
	}

	expense.ID = ID

	json.NewEncoder(w).Encode(expense)
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
		}).Error("error parsing expense id")

		// TODO: Return an error response
	}

	expense := models.UserExpense{}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error("error decoding expense payload")

		// TODO: Return an error response
	}

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error("error updating expense")

		// TODO: Return an error response
	}

	if count == 0 {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.UpdateExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Warn("update expense modified 0 rows")

		// TODO: Return an error response
		json.NewEncoder(w).Encode(&models.UserExpense{})

		return
	}

	expense.ID = ID

	json.NewEncoder(w).Encode(expense)
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
		}).Error("error parsing expense id")

		// TODO: Return an error response
	}

	count, err := e.service.DeleteExpense(ID, userAccountID)
	if err != nil {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Error("error deleting expense")

		// TODO: Return an error response
	}

	if count == 0 {
		log.WithFields(log.Fields{
			"method":     "ExpenseController.DeleteExpense",
			"account ID": userAccountID,
			"id":         ID,
			"err":        err.Error(),
		}).Warn("delete expense removed 0 rows")

		// TODO: Return an error response
		json.NewEncoder(w).Encode(&models.UserExpense{})

		return
	}

	json.NewEncoder(w).Encode(ID)
}
