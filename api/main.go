package main

import (
	"net/http"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	typesController := controllers.NewTypesController()
	categoriesController := controllers.NewCategoriesController()
	expensesController := controllers.NewExpensesController()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	// TODO: TBD

	// Types
	api.HandleFunc("/types", typesController.GetTypes).Methods("GET")

	// Categories
	api.HandleFunc("/categories", categoriesController.GetCategories).Methods("GET")

	// Expenses
	api.HandleFunc("/expenses", expensesController.GetExpenses).Methods("GET")
	api.HandleFunc("/expenses", expensesController.CreateExpense).Methods("POST")
	api.HandleFunc("/expenses/{id}", expensesController.UpdateExpense).Methods("PUT")
	api.HandleFunc("/expenses/{id}", expensesController.DeleteExpense).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
