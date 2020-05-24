package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

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

func NewExpenseService(logger *logrus.Logger, repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{
		logger,
		repo,
	}
}

func (e *expenseService) GetExpenses(userAccountID int64) ([]models.Expense, error) {
	return e.repo.GetExpenses(userAccountID)
}

func (e *expenseService) CreateExpense(userAccountID int64, expense *models.Expense) (int64, error) {
	return e.repo.CreateExpense(userAccountID, expense)
}

func (e *expenseService) UpdateExpense(ID, userAccountID int64, expense *models.Expense) (int64, error) {
	return e.repo.UpdateExpense(ID, userAccountID, expense)
}

func (e *expenseService) DeleteExpense(ID, userAccountID int64) (int64, error) {
	return e.repo.DeleteExpense(ID, userAccountID)
}
