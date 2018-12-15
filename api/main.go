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
	// DB connection to the MySQL instance
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
	authRepository := repositories.NewAuthRepository(DB)
	typeRepository := repositories.NewTypeRepository(DB)
	categoryRepository := repositories.NewCategoryRepository(DB)
	expenseRepository := repositories.NewExpenseRepository(DB)

	authService := services.NewAuthService(authRepository)
	typeService := services.NewTypeService(typeRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	expenseService := services.NewExpenseService(expenseRepository)

	authController := controllers.NewAuthController(authService)
	typeController := controllers.NewTypeController(typeService)
	categoryController := controllers.NewCategoryController(categoryService)
	expenseController := controllers.NewExpenseController(expenseService)

	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()

	// Auth
	api.HandleFunc("/accounts", authController.CreateAccount).Methods("POST")

	// Types
	api.HandleFunc("/types", typeController.GetTypes).Methods("GET")

	// Categories
	api.HandleFunc("/categories", categoryController.GetCategories).Methods("GET")

	// Expenses
	api.HandleFunc("/expenses", expenseController.GetExpenses).Methods("GET")
	api.HandleFunc("/expenses", expenseController.CreateExpense).Methods("POST")
	api.HandleFunc("/expenses/{id}", expenseController.UpdateExpense).Methods("PUT")
	api.HandleFunc("/expenses/{id}", expenseController.DeleteExpense).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
