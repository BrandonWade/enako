package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_expense_repository.go . ExpenseRepository
type ExpenseRepository interface {
	GetExpenses(userAccountID int64) ([]models.Expense, error)
	CreateExpense(userAccountID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, userAccountID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, userAccountID int64) (int64, error)
}

type expenseRepository struct {
	DB *sqlx.DB
}

func NewExpenseRepository(DB *sqlx.DB) ExpenseRepository {
	return &expenseRepository{
		DB,
	}
}

func (e *expenseRepository) GetExpenses(userAccountID int64) ([]models.Expense, error) {
	expenses := []models.Expense{}

	err := e.DB.Select(&expenses, `SELECT
		e.id,
		c.name category,
		e.description,
		e.amount / 100 amount,
		DATE(expense_date) expense_date
        FROM expenses e
		INNER JOIN categories c ON c.id = e.category_id
        WHERE e.user_account_id = ?;
    `, userAccountID)
	if err != nil {
		return []models.Expense{}, err
	}

	return expenses, nil
}

func (e *expenseRepository) CreateExpense(userAccountID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`INSERT
		INTO expenses(
			user_account_id,
			category_id,
			description,
			amount,
			expense_date
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?
		);
	`,
		userAccountID,
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

func (e *expenseRepository) UpdateExpense(ID, userAccountID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`UPDATE expenses
		SET
			user_account_id = ?,
			category_id = ?,
			description = ?,
			amount = ?,
			date = ?
		WHERE id = ?;
	`,
		userAccountID,
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

func (e *expenseRepository) DeleteExpense(ID, userAccountID int64) (int64, error) {
	result, err := e.DB.Exec(`DELETE
		FROM expenses
		WHERE id = ?
		AND user_account_id = ?;
	`,
		ID,
		userAccountID,
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
