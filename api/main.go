package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth
	// TODO: TBD

	// Types
	api.HandleFunc("/types", GetTypes).Methods("GET")

	// Categories
	api.HandleFunc("/categories", GetCategories).Methods("GET")

	// Expenses
	api.HandleFunc("/expenses", GetExpenses).Methods("GET")
	api.HandleFunc("/expenses", CreateExpense).Methods("POST")
	api.HandleFunc("/expenses/{id}", UpdateExpense).Methods("PUT")
	api.HandleFunc("/expenses/{id}", DeleteExpense).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func GetTypes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get types"))
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get categories"))
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get expenses"))
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create expense"))
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update expense"))
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete expense"))
}
