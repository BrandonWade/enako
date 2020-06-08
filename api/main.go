package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/controllers/middleware"
	"github.com/BrandonWade/enako/api/helpers"
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

	cookieSecret string
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

	cookieSecret = os.Getenv("COOKIE_SECRET")

	logger = logrus.New()

	validation.InitValidator()
}

func main() {
	hasher := helpers.NewPasswordHasher(logger)
	store := helpers.NewCookieStore([]byte(cookieSecret))

	stack := middleware.NewMiddlewareStack(logger, store)

	authRepository := repositories.NewAuthRepository(DB)
	categoryRepository := repositories.NewCategoryRepository(DB)
	expenseRepository := repositories.NewExpenseRepository(DB)

	authService := services.NewAuthService(logger, hasher, authRepository)
	categoryService := services.NewCategoryService(logger, categoryRepository)
	expenseService := services.NewExpenseService(logger, expenseRepository)

	authController := controllers.NewAuthController(logger, store, authService)
	categoryController := controllers.NewCategoryController(logger, store, categoryService)
	expenseController := controllers.NewExpenseController(logger, store, expenseService)

	// Set up route middleware
	getCategoriesHandler := stack.Apply(categoryController.GetCategories, []middleware.Middleware{stack.Authenticate()})
	getExpensesHandler := stack.Apply(expenseController.GetExpenses, []middleware.Middleware{stack.Authenticate()})
	createExpenseHandler := stack.Apply(expenseController.CreateExpense, []middleware.Middleware{stack.Authenticate(), stack.DecodeExpense()})
	updateExpenseHandler := stack.Apply(expenseController.UpdateExpense, []middleware.Middleware{stack.Authenticate(), stack.DecodeExpense()})
	deleteExpenseHandler := stack.Apply(expenseController.DeleteExpense, []middleware.Middleware{stack.Authenticate()})

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()

	// Auth
	api.HandleFunc("/accounts", authController.CreateAccount).Methods("POST")
	api.HandleFunc("/login", authController.Login).Methods("POST")
	api.HandleFunc("/logout", authController.Logout).Methods("GET")

	// Categories
	api.HandleFunc("/categories", getCategoriesHandler).Methods("GET")

	// Expenses
	api.HandleFunc("/expenses", getExpensesHandler).Methods("GET")
	api.HandleFunc("/expenses", createExpenseHandler).Methods("POST")
	api.HandleFunc("/expenses/{id}", updateExpenseHandler).Methods("PUT")
	api.HandleFunc("/expenses/{id}", deleteExpenseHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
