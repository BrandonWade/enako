package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

// ExpenseService an interface for working with expenses.
//go:generate counterfeiter -o fakes/fake_expense_service.go . ExpenseService
type ExpenseService interface {
	GetExpenses(userAccountID int64) (expenses []models.Expense, err error)
	CreateExpense(userAccountID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, userAccountID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, userAccountID int64) (int64, error)
}

type expenseService struct {
	logger *logrus.Logger
	repo   repositories.ExpenseRepository
}

// NewExpenseService returns a new instance of an ExpenseService.
func NewExpenseService(logger *logrus.Logger, repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{
		logger,
		repo,
	}
}

// GetExpenses retrieves the expenses belonging to the given user account id.
func (e *expenseService) GetExpenses(userAccountID int64) ([]models.Expense, error) {
	return e.repo.GetExpenses(userAccountID)
}

// CreateExpense creates an expense belonging to the given user account id.
func (e *expenseService) CreateExpense(userAccountID int64, expense *models.Expense) (int64, error) {
	return e.repo.CreateExpense(userAccountID, expense)
}

// DeleteExpense updates the expense with the given ID belonging to the given user account id.
func (e *expenseService) UpdateExpense(ID, userAccountID int64, expense *models.Expense) (int64, error) {
	return e.repo.UpdateExpense(ID, userAccountID, expense)
}

// DeleteExpense deletes the expense with the given ID belonging to the given user account id.
func (e *expenseService) DeleteExpense(ID, userAccountID int64) (int64, error) {
	return e.repo.DeleteExpense(ID, userAccountID)
}
