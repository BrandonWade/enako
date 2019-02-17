package models

type ExpenseType struct {
	ID        int64  `json:"id" db:"id"`
	TypeName  string `json:"type_name,omitempty" db:"type_name"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

type ExpenseCategory struct {
	ID           int64  `json:"id" db:"id"`
	CategoryName string `json:"category_name,omitempty" db:"category_name"`
	CreatedAt    string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    string `json:"updated_at,omitempty" db:"updated_at"`
}

type UserAccount struct {
	ID                  int64  `json:"id" db:"id"`
	UserAccountEmail    string `json:"user_account_email" db:"user_account_email" validate:"email"`
	UserAccountPassword string `json:"user_account_password,omitempty" db:"user_account_password" validate:"pword"`
	ConfirmPassword     string `json:"confirm_password,omitempty" validate:"pword"`
}

type UserExpense struct {
	ID                 int64  `json:"id" db:"id"`
	UserAccountID      int64  `json:"user_account_id,omitempty" db:"user_account_id"`
	ExpenseType        string `json:"expense_type,omitempty" db:"expense_type"`
	ExpenseTypeID      int64  `json:"expense_type_id,omitempty" db:"expense_type_id" validate:"min=1"`
	ExpenseCategory    string `json:"expense_category,omitempty" db:"expense_category"`
	ExpenseCategoryID  int64  `json:"expense_category_id,omitempty" db:"expense_category_id" validate:"min=1"`
	ExpenseDescription string `json:"expense_description,omitempty" db:"expense_description"`
	ExpenseAmount      int64  `json:"expense_amount,omitempty" db:"expense_amount" validate:"min=1"`
	ExpenseDate        string `json:"expense_date,omitempty" db:"expense_date" validate:"date"`
	CreatedAt          string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          string `json:"updated_at,omitempty" db:"updated_at"`
}
