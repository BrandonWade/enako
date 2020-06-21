package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

// ExpenseService an interface for working with expenses.
//go:generate counterfeiter -o fakes/fake_expense_service.go . ExpenseService
type ExpenseService interface {
	GetExpenses(accountID int64) (expenses []models.Expense, err error)
	CreateExpense(accountID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, accountID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, accountID int64) (int64, error)
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

// GetExpenses retrieves the expenses belonging to the given account id.
func (e *expenseService) GetExpenses(accountID int64) ([]models.Expense, error) {
	return e.repo.GetExpenses(accountID)
}

// CreateExpense creates an expense belonging to the given account id.
func (e *expenseService) CreateExpense(accountID int64, expense *models.Expense) (int64, error) {
	return e.repo.CreateExpense(accountID, expense)
}

// DeleteExpense updates the expense with the given ID belonging to the given account id.
func (e *expenseService) UpdateExpense(ID, accountID int64, expense *models.Expense) (int64, error) {
	return e.repo.UpdateExpense(ID, accountID, expense)
}

// DeleteExpense deletes the expense with the given ID belonging to the given account id.
func (e *expenseService) DeleteExpense(ID, accountID int64) (int64, error) {
	return e.repo.DeleteExpense(ID, accountID)
}
