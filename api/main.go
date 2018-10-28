package main

import (
	"net/http"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	expenseController := controllers.NewExpenseController()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	// TODO: TBD

	// Types
	api.HandleFunc("/types", GetTypes).Methods("GET")

	// Categories
	api.HandleFunc("/categories", GetCategories).Methods("GET")

	// Expenses
	api.HandleFunc("/expenses", expenseController.GetExpenses).Methods("GET")
	api.HandleFunc("/expenses", expenseController.CreateExpense).Methods("POST")
	api.HandleFunc("/expenses/{id}", expenseController.UpdateExpense).Methods("PUT")
	api.HandleFunc("/expenses/{id}", expenseController.DeleteExpense).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func GetTypes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get types"))
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get categories"))
}
