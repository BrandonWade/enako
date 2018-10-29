package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type ExpensesService interface {
	GetExpenses() (expenses []models.Expense, err error)
	CreateExpense(userID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, userID int64, expense *models.Expense) error
}

type expensesService struct {
	repo repositories.ExpensesRepository
}

func NewExpensesService(repo repositories.ExpensesRepository) ExpensesService {
	return &expensesService{
		repo,
	}
}

func (e *expensesService) GetExpenses() ([]models.Expense, error) {
	return e.repo.GetExpenses()
}

func (e *expensesService) CreateExpense(userID int64, expense *models.Expense) (int64, error) {
	return e.repo.CreateExpense(userID, expense)
}

func (e *expensesService) UpdateExpense(ID, userID int64, expense *models.Expense) error {
	return e.repo.UpdateExpense(ID, userID, expense)
}
