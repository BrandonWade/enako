package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/BrandonWade/enako/api/services"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func init() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting to db: %s\n", err.Error())
	}
}

func main() {
	typesRepository := repositories.NewTypesRepository(DB)
	categoriesRepository := repositories.NewCategoriesRepository(DB)
	expensesRepository := repositories.NewExpensesRepository(DB)

	typesService := services.NewTypesService(typesRepository)
	categoriesService := services.NewCategoriesService(categoriesRepository)
	expensesService := services.NewExpensesService(expensesRepository)

	typesController := controllers.NewTypesController(typesService)
	categoriesController := controllers.NewCategoriesController(categoriesService)
	expensesController := controllers.NewExpensesController(expensesService)

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
