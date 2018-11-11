package models

type ExpenseType struct {
	ID              int64  `json:"id" db:"id"`
	ExpenseTypeName string `json:"expense_type_name,omitempty" db:"expense_type_name"`
	CreatedAt       string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       string `json:"updated_at,omitempty" db:"updated_at"`
}

type ExpenseCategory struct {
	ID                  int64  `json:"id" db:"id"`
	ExpenseCategoryName string `json:"expense_category_name,omitempty" db:"expense_category_name"`
	CreatedAt           string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt           string `json:"updated_at,omitempty" db:"updated_at"`
}

type UserAccount struct {
	ID                  int64  `json:"id" db:"id"`
	UserAccountEmail    string `json:"user_account_email" db:"user_account_email"`
	UserAccountPassword string `json:"user_account_password" db:"user_account_password"`
	ConfirmPassword     string `json:"confirm_password"`
}

type UserExpense struct {
	ID                 int64  `json:"id" db:"id"`
	UserAccountID      int64  `json:"user_account_id,omitempty" db:"user_account_id"`
	ExpenseType        string `json:"expense_type,omitempty" db:"expense_type"`
	ExpenseCategory    string `json:"expense_category,omitempty" db:"expense_category"`
	ExpenseDescription string `json:"expense_description,omitempty" db:"expense_description"`
	ExpenseAmount      int64  `json:"expense_amount,omitempty" db:"expense_amount"`
	ExpenseDate        string `json:"expense_date,omitempty" db:"expense_date"`
	CreatedAt          string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          string `json:"updated_at,omitempty" db:"updated_at"`
}
