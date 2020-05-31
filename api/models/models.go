package models

type Category struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name,omitempty" db:"name"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

type UserAccount struct {
	ID              int64  `json:"id" db:"id"`
	Username        string `json:"username" db:"username" validate:"uname"`
	Email           string `json:"email" db:"email" validate:"email"`
	Password        string `json:"password,omitempty" db:"password" validate:"pword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"pword"`
}

type Expense struct {
	ID            int64   `json:"id" db:"id"`
	UserAccountID int64   `json:"user_account_id,omitempty" db:"user_account_id"`
	Category      string  `json:"category,omitempty" db:"category"`
	CategoryID    int64   `json:"category_id,omitempty" db:"category_id" validate:"min=1"`
	Description   string  `json:"description,omitempty" db:"description"`
	Amount        float64 `json:"amount,omitempty" db:"amount" validate:"min=1"`
	ExpenseDate   string  `json:"expense_date,omitempty" db:"expense_date" validate:"date"`
	CreatedAt     string  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     string  `json:"updated_at,omitempty" db:"updated_at"`
}
