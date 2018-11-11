package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type ExpenseRepository interface {
	GetExpenses() ([]models.UserExpense, error)
	CreateExpense(userAccountID int64, expense *models.UserExpense) (int64, error)
	UpdateExpense(ID, userAccountID int64, expense *models.UserExpense) (int64, error)
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

func (e *expenseRepository) GetExpenses() ([]models.UserExpense, error) {
	userAccountID := 1 // TODO: Hardcoded for testing
	expenses := []models.UserExpense{}

	err := e.DB.Select(&expenses, `SELECT *
        FROM user_expenses AS e
        WHERE e.user_account_id = ?;
    `, userAccountID)
	if err != nil {
		return []models.UserExpense{}, err
	}

	return expenses, nil
}

func (e *expenseRepository) CreateExpense(userAccountID int64, expense *models.UserExpense) (int64, error) {
	result, err := e.DB.Exec(`INSERT
		INTO user_expenses(
			user_account_id,
			expense_type,
			expense_category,
			expense_description,
			expense_amount,
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
		expense.ExpenseType,
		expense.ExpenseCategory,
		expense.ExpenseDescription,
		expense.ExpenseAmount,
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

func (e *expenseRepository) UpdateExpense(ID, userAccountID int64, expense *models.UserExpense) (int64, error) {
	result, err := e.DB.Exec(`UPDATE user_expenses
		SET
			user_account_id = ?,
			expense_type = ?,
			expense_category = ?,
			expense_description = ?,
			expense_amount = ?,
			expense_date = ?
		WHERE id = ?;
	`,
		userAccountID,
		expense.ExpenseType,
		expense.ExpenseCategory,
		expense.ExpenseDescription,
		expense.ExpenseAmount,
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
		FROM user_expenses
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
