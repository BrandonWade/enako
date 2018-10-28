package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/gorilla/mux"
)

type ExpenseController interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
}

type expenseController struct {
	expenses []models.Expense
}

func NewExpenseController() ExpenseController {
	return &expenseController{
		[]models.Expense{ // TODO: Hardcoded for testing
			models.Expense{
				1,
				"unnecessary",
				"food",
				"went out for lunch",
				1680,
				"October 15th 2018",
			},
			models.Expense{
				2,
				"recurring",
				"technology",
				"paid phone bill for next 2 months",
				12058,
				"October 16th 2018",
			},
			models.Expense{
				3,
				"unnecessary",
				"entertainment",
				"went to a movie",
				1150,
				"October 17th 2018",
			},
		},
	}
}

func (e *expenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(e.expenses)
}

func (e *expenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var newExpense models.Expense
	err := json.NewDecoder(r.Body).Decode(&newExpense)
	if err != nil {
		// TODO: Handle
	}

	ID := 1
	if len(e.expenses) > 0 {
		ID = e.expenses[len(e.expenses)-1].ID + 1
	}

	newExpense.ID = ID
	e.expenses = append(e.expenses, newExpense)

	json.NewEncoder(w).Encode(e.expenses)
}

func (e *expenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: Handle
	}

	var newExpense models.Expense
	err = json.NewDecoder(r.Body).Decode(&newExpense)
	if err != nil {
		// TODO: Handle
	}

	for i, expense := range e.expenses {
		if expense.ID == ID {
			e.expenses[i] = newExpense
			break
		}
	}

	json.NewEncoder(w).Encode(e.expenses)
}

func (e *expenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: Handle
	}

	for i, expense := range e.expenses {
		if expense.ID == ID {
			e.expenses = append(e.expenses[:i], e.expenses[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(e.expenses)
}
