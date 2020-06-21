package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// ExpenseRepository an interface for working with expenses.
//go:generate counterfeiter -o fakes/fake_expense_repository.go . ExpenseRepository
type ExpenseRepository interface {
	GetExpenses(accountID int64) ([]models.Expense, error)
	CreateExpense(accountID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, accountID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, accountID int64) (int64, error)
}

type expenseRepository struct {
	DB *sqlx.DB
}

// NewExpenseRepository returns a new instance of an ExpenseRepository.
func NewExpenseRepository(DB *sqlx.DB) ExpenseRepository {
	return &expenseRepository{
		DB,
	}
}

// GetExpenses retrieves the expenses belonging to the given account ID.
func (e *expenseRepository) GetExpenses(accountID int64) ([]models.Expense, error) {
	expenses := []models.Expense{}

	err := e.DB.Select(&expenses, `SELECT
		e.id,
		c.id category_id,
		c.name category,
		e.description,
		e.amount / 100 amount,
		DATE(expense_date) expense_date
        FROM expenses e
		INNER JOIN categories c ON c.id = e.category_id
        WHERE e.account_id = ?;
    `, accountID)
	if err != nil {
		return []models.Expense{}, err
	}

	return expenses, nil
}

// CreateExpense creates an expense belonging to the given account ID.
func (e *expenseRepository) CreateExpense(accountID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`INSERT
		INTO expenses(
			account_id,
			category_id,
			description,
			amount,
			expense_date
		) VALUES (
			?,
			?,
			?,
			?,
			?
		);
	`,
		accountID,
		expense.CategoryID,
		expense.Description,
		expense.Amount,
		expense.ExpenseDate,
	)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID, nil
}

// UpdateExpense updates the expense with the given ID belonging to the given account ID.
func (e *expenseRepository) UpdateExpense(ID, accountID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`UPDATE expenses
		SET
			account_id = ?,
			category_id = ?,
			description = ?,
			amount = ?,
			expense_date = ?
		WHERE id = ?;
	`,
		accountID,
		expense.CategoryID,
		expense.Description,
		expense.Amount,
		expense.ExpenseDate,
		ID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// DeleteExpense deletes the expense with the given ID belonging to the given account ID.
func (e *expenseRepository) DeleteExpense(ID, accountID int64) (int64, error) {
	result, err := e.DB.Exec(`DELETE
		FROM expenses
		WHERE id = ?
		AND account_id = ?;
	`,
		ID,
		accountID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
