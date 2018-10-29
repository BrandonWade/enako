package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"
)

type ExpensesController interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
}

type expensesController struct {
	service services.ExpensesService
}

func NewExpensesController(service services.ExpensesService) ExpensesController {
	return &expensesController{
		service,
	}
}

func (e *expensesController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := e.service.GetExpenses()
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(expenses)
}

func (e *expensesController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userID := int64(1) // TODO: Hardcoded for testing

	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		// TODO: Handle
	}

	ID, err := e.service.CreateExpense(userID, &expense)
	if err != nil {
		// TODO: Handle
	}

	// TODO: Get from DB
	expense.ID = ID

	json.NewEncoder(w).Encode(expense)
}

func (e *expensesController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userID := int64(1) // TODO: Hardcoded for testing

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		// TODO: Handle
	}

	expense := models.Expense{}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		// TODO: Handle
	}

	count, err := e.service.UpdateExpense(ID, userID, &expense)
	if err != nil {
		// TODO: Handle
	}

	if count == 0 {
		json.NewEncoder(w).Encode(&models.Expense{})
		return
	}

	// TODO: Get from DB
	expense.ID = ID

	json.NewEncoder(w).Encode(expense)
}

func (e *expensesController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	userID := int64(1) // TODO: Hardcoded for testing

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		// TODO: Handle
	}

	count, err := e.service.DeleteExpense(ID, userID)
	if err != nil {
		// TODO: Handle
	}

	if count == 0 {
		json.NewEncoder(w).Encode(&models.Expense{})
		return
	}

	json.NewEncoder(w).Encode(ID)
}
