package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BrandonWade/enako/api/clients"
	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/BrandonWade/enako/api/services"
	"github.com/BrandonWade/enako/api/validation"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

var (
	// DB connection to the MySQL instance
	DB *sqlx.DB

	logger *logrus.Logger

	mjClient *mailjet.Client

	basePath     string
	cookieSecret string
	csrfSecret   string
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

	mjPublicKey := os.Getenv("MAILJET_PUBLIC_KEY")
	mjPrivateKey := os.Getenv("MAILJET_PRIVATE_KEY")
	mjClient = mailjet.NewMailjetClient(mjPublicKey, mjPrivateKey)

	cookieSecret = os.Getenv("COOKIE_SECRET")
	csrfSecret = os.Getenv("CSRF_SECRET")

	basePath, err = os.Getwd()
	if err != nil {
		log.Fatalf("error fetching base path: %s\n", err.Error())
	}

	validation.InitValidator()
}

func main() {
	csrfMiddleware := csrf.Protect([]byte(csrfSecret))
	hasher := helpers.NewPasswordHasher(logger)
	templater := helpers.NewTemplater(logger)
	generator := helpers.NewTokenGenerator()
	store := helpers.NewCookieStore([]byte(cookieSecret))

	stack := middleware.NewMiddlewareStack(logger, store)

	mailjetClient := clients.NewMailjetClient(logger, mjClient)

	templateService := services.NewTemplateService(logger, templater, basePath)
	emailService := services.NewEmailService(logger, templateService, mailjetClient)

	accountRepository := repositories.NewAccountRepository(DB)
	passwordResetRepository := repositories.NewPasswordResetRepository(DB)
	categoryRepository := repositories.NewCategoryRepository(DB)
	expenseRepository := repositories.NewExpenseRepository(DB)

	accountService := services.NewAccountService(logger, hasher, generator, emailService, accountRepository)
	passwordResetService := services.NewPasswordResetService(logger, hasher, generator, emailService, passwordResetRepository, accountRepository)
	categoryService := services.NewCategoryService(logger, categoryRepository)
	expenseService := services.NewExpenseService(logger, expenseRepository)

	authController := controllers.NewAuthController(logger, store, accountService)
	passwordResetController := controllers.NewPasswordResetController(logger, passwordResetService)
	accountController := controllers.NewAccountController(logger, accountService)
	categoryController := controllers.NewCategoryController(logger, store, categoryService)
	expenseController := controllers.NewExpenseController(logger, store, expenseService)

	// Set up route middleware
	loginHandler := middleware.Apply(authController.Login, []middleware.Middleware{stack.DecodeLogin()})

	registerUserHandler := middleware.Apply(accountController.RegisterUser, []middleware.Middleware{stack.ValidateCreateAccount(), stack.DecodeCreateAccount()})
	changePasswordHandler := middleware.Apply(accountController.ChangePassword, []middleware.Middleware{stack.ValidateChangePassword(), stack.DecodeChangePassword(), stack.Authenticate()})

	requestPasswordResetHander := middleware.Apply(passwordResetController.RequestPasswordReset, []middleware.Middleware{stack.ValidateRequestPasswordReset(), stack.DecodeRequestPasswordReset()})
	passwordResetHander := middleware.Apply(passwordResetController.ResetPassword, []middleware.Middleware{stack.ValidatePasswordReset(), stack.DecodePasswordReset()})

	getCategoriesHandler := middleware.Apply(categoryController.GetCategories, []middleware.Middleware{stack.Authenticate()})

	getExpensesHandler := middleware.Apply(expenseController.GetExpenses, []middleware.Middleware{stack.Authenticate()})
	createExpenseHandler := middleware.Apply(expenseController.CreateExpense, []middleware.Middleware{stack.ValidateExpense(), stack.DecodeExpense(), stack.Authenticate()})
	updateExpenseHandler := middleware.Apply(expenseController.UpdateExpense, []middleware.Middleware{stack.ValidateExpense(), stack.DecodeExpense(), stack.Authenticate()})
	deleteExpenseHandler := middleware.Apply(expenseController.DeleteExpense, []middleware.Middleware{stack.Authenticate()})

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()

	// Auth
	authAPI := api.PathPrefix("").Subrouter()
	authAPI.Use(csrfMiddleware)
	authAPI.HandleFunc("/csrf", authController.CSRF).Methods("HEAD")
	authAPI.HandleFunc("/login", loginHandler).Methods("POST")
	authAPI.HandleFunc("/logout", authController.Logout).Methods("GET")

	// Accounts
	accountAPI := api.PathPrefix("/accounts").Subrouter()
	accountAPI.Use(csrfMiddleware)
	accountAPI.HandleFunc("", registerUserHandler).Methods("POST")
	accountAPI.HandleFunc("/activate", accountController.ActivateAccount).Methods("GET")
	accountAPI.HandleFunc("/password", requestPasswordResetHander).Methods("POST")
	accountAPI.HandleFunc("/password/reset", passwordResetController.SetPasswordResetToken).Methods("GET")
	accountAPI.HandleFunc("/password/reset", passwordResetHander).Methods("POST")
	accountAPI.HandleFunc("/password/change", changePasswordHandler).Methods("PUT")

	// Categories
	categoryAPI := api.PathPrefix("/categories").Subrouter()
	categoryAPI.Use(csrfMiddleware)
	categoryAPI.HandleFunc("", getCategoriesHandler).Methods("GET")

	// Expenses
	expenseAPI := api.PathPrefix("/expenses").Subrouter()
	expenseAPI.Use(csrfMiddleware)
	expenseAPI.HandleFunc("", getExpensesHandler).Methods("GET")
	expenseAPI.HandleFunc("", createExpenseHandler).Methods("POST")
	expenseAPI.HandleFunc("/{id}", updateExpenseHandler).Methods("PUT")
	expenseAPI.HandleFunc("/{id}", deleteExpenseHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
