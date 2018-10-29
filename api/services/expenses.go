package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type ExpenseService interface {
	GetExpenses() (expenses []models.Expense, err error)
	CreateExpense(userID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, userID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, userID int64) (int64, error)
}

type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{
		repo,
	}
}

func (e *expenseService) GetExpenses() ([]models.Expense, error) {
	return e.repo.GetExpenses()
}

func (e *expenseService) CreateExpense(userID int64, expense *models.Expense) (int64, error) {
	return e.repo.CreateExpense(userID, expense)
}

func (e *expenseService) UpdateExpense(ID, userID int64, expense *models.Expense) (int64, error) {
	return e.repo.UpdateExpense(ID, userID, expense)
}

func (e *expenseService) DeleteExpense(ID, userID int64) (int64, error) {
	return e.repo.DeleteExpense(ID, userID)
}
