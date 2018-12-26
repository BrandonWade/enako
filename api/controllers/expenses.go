package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"
)

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

func NewExpenseController(service services.ExpenseService) ExpenseController {
	return &expenseController{
		service,
	}
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	expenses, err := e.service.GetExpenses(userAccountID)
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(expenses)
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	var expense models.UserExpense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		// TODO: Handle
	}

	ID, err := e.service.CreateExpense(userAccountID, &expense)
	if err != nil {
		// TODO: Handle
	}

	expense.ID = ID

	json.NewEncoder(w).Encode(expense)
}

func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userAccountID := int64(1) // TODO: Hardcoded for testing

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		// TODO: Handle
	}

	expense := models.UserExpense{}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		// TODO: Handle
	}

	count, err := e.service.UpdateExpense(ID, userAccountID, &expense)
	if err != nil {
		// TODO: Handle
	}

	if count == 0 {
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
		// TODO: Handle
	}

	count, err := e.service.DeleteExpense(ID, userAccountID)
	if err != nil {
		// TODO: Handle
	}

	if count == 0 {
		json.NewEncoder(w).Encode(&models.UserExpense{})
		return
	}

	json.NewEncoder(w).Encode(ID)
}
