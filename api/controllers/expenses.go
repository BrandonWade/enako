package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrandonWade/enako/api/models"
	"github.com/gorilla/mux"
)

type ExpensesController interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
	DeleteExpense(w http.ResponseWriter, r *http.Request)
}

type expensesController struct {
	expenses []models.Expense
}

func NewExpensesController() ExpensesController {
	return &expensesController{
		[]models.Expense{ // TODO: Hardcoded for testing
			models.Expense{
				ID:          1,
				Type:        "unnecessary",
				Category:    "food",
				Description: "went out for lunch",
				Amount:      1680,
				Date:        "October 15th 2018",
			},
			models.Expense{
				ID:          2,
				Type:        "recurring",
				Category:    "technology",
				Description: "paid phone bill for next 2 months",
				Amount:      12058,
				Date:        "October 16th 2018",
			},
			models.Expense{
				ID:          3,
				Type:        "unnecessary",
				Category:    "entertainment",
				Description: "went to a movie",
				Amount:      1150,
				Date:        "October 17th 2018",
			},
		},
	}
}

func (e *expensesController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(e.expenses)
}

func (e *expensesController) CreateExpense(w http.ResponseWriter, r *http.Request) {
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

func (e *expensesController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
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

func (e *expensesController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
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
