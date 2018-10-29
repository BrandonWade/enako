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
	expenses []models.Expense
	service  services.ExpensesService
}

func NewExpensesController(service services.ExpensesService) ExpensesController {
	return &expensesController{
		expenses: []models.Expense{ // TODO: Hardcoded for testing
			models.Expense{
				ID:          4,
				Type:        "unnecessary",
				Category:    "food",
				Description: "went out for lunch",
				Amount:      1680,
				Date:        "October 15th 2018",
			},
			models.Expense{
				ID:          5,
				Type:        "recurring",
				Category:    "technology",
				Description: "paid phone bill for next 2 months",
				Amount:      12058,
				Date:        "October 16th 2018",
			},
			models.Expense{
				ID:          6,
				Type:        "unnecessary",
				Category:    "entertainment",
				Description: "went to a movie",
				Amount:      1150,
				Date:        "October 17th 2018",
			},
		},
		service: service,
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
	expense.UserID = userID

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

	err = e.service.UpdateExpense(ID, userID, &expense)
	if err != nil {
		// TODO: Handle
	}

	// TODO: Get from DB
	expense.ID = ID
	expense.UserID = userID

	json.NewEncoder(w).Encode(expense)
}

func (e *expensesController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		// TODO: Handle
	}

	for i, expense := range e.expenses {
		if expense.ID == int64(ID) {
			e.expenses = append(e.expenses[:i], e.expenses[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(e.expenses)
}
