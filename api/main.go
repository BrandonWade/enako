package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/BrandonWade/enako/api/services"
	"github.com/BrandonWade/enako/api/validation"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"
)

var (
	// DB connection to the MySQL instance
	DB *sqlx.DB

	logger *logrus.Logger
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

	logger = logrus.New()

	validation.InitValidator()
}

func main() {
	authRepository := repositories.NewAuthRepository(DB)
	typeRepository := repositories.NewTypeRepository(DB)
	categoryRepository := repositories.NewCategoryRepository(DB)
	expenseRepository := repositories.NewExpenseRepository(DB)

	authService := services.NewAuthService(logger, authRepository)
	typeService := services.NewTypeService(logger, typeRepository)
	categoryService := services.NewCategoryService(logger, categoryRepository)
	expenseService := services.NewExpenseService(logger, expenseRepository)

	authController := controllers.NewAuthController(logger, authService)
	typeController := controllers.NewTypeController(logger, typeService)
	categoryController := controllers.NewCategoryController(logger, categoryService)
	expenseController := controllers.NewExpenseController(logger, expenseService)

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
