package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_expense_repository.go . ExpenseRepository
type ExpenseRepository interface {
	GetExpenses(userAccountID int64) ([]models.UserExpense, error)
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

func (e *expenseRepository) GetExpenses(userAccountID int64) ([]models.UserExpense, error) {
	expenses := []models.UserExpense{}

	err := e.DB.Select(&expenses, `SELECT
		e.id,
		t.type_name AS expense_type,
		c.category_name AS expense_category,
		e.expense_description,
		e.expense_amount / 100 AS expense_amount,
		DATE(expense_date) AS expense_date
        FROM user_expenses AS e
		INNER JOIN expense_types AS t ON t.id = e.expense_type_id
		INNER JOIN expense_categories AS c ON c.id = e.expense_category_id
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
			expense_type_id,
			expense_category_id,
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
		expense.ExpenseTypeID,
		expense.ExpenseCategoryID,
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
			expense_type_id = ?,
			expense_category_id = ?,
			expense_description = ?,
			expense_amount = ?,
			expense_date = ?
		WHERE id = ?;
	`,
		userAccountID,
		expense.ExpenseTypeID,
		expense.ExpenseCategoryID,
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
