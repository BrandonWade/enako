package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"regexp"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/BrandonWade/enako/api/services"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
	validator "gopkg.in/validator.v2"
)

const (
	minPasswordLength = 15
	maxPasswordLength = 50
)

var (
	// DB connection to the MySQL instance
	DB *sqlx.DB

	errMustBeString    = errors.New("must be string")
	errInvalidEmail    = errors.New("invalid email")
	errInvalidPassword = errors.New("invalid password")
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

	// Add a simple email validation rule
	validator.SetValidationFunc("email", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		match, err := regexp.MatchString("^[^@]+@[^\\.@]+\\..+$", t.String())
		if err != nil || match != true {
			return errInvalidEmail
		}

		return nil
	})

	// Add a password matching rule (alphanumeric plus symbols)
	validator.SetValidationFunc("pword", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		pword := t.String()

		// Ensure length is compatible with bcrypt requirements
		if len(pword) < 15 || len(pword) > 50 {
			return errInvalidPassword
		}

		match, err := regexp.MatchString("^[\\w\\!\\@\\#\\$\\%\\^\\&\\*]+$", pword)
		if err != nil || match != true {
			return errInvalidPassword
		}

		return nil
	})
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
