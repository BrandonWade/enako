package models

// Category a model for working with expense categories.
type Category struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name,omitempty" db:"name"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

// Account a model for working with accounts.
type Account struct {
	ID          int64  `json:"id" db:"id"`
	Username    string `json:"username" db:"username" validate:"uname"`
	Email       string `json:"email,omitempty" db:"email" validate:"email"`
	Password    string `json:"password,omitempty" db:"password" validate:"pword"`
	IsActivated bool   `json:"is_activated,omitempty" db:"is_activated"`
	CreatedAt   string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty" db:"updated_at"`
}

// Expense a model for working with expenses.
type Expense struct {
	ID          int64   `json:"id" db:"id"`
	AccountID   int64   `json:"account_id,omitempty" db:"account_id"`
	Category    string  `json:"category,omitempty" db:"category"`
	CategoryID  int64   `json:"category_id,omitempty" db:"category_id" validate:"min=1"`
	Description string  `json:"description,omitempty" db:"description" validate:"nonzero"`
	Amount      float64 `json:"amount,omitempty" db:"amount" validate:"min=1"`
	ExpenseDate string  `json:"expense_date,omitempty" db:"expense_date" validate:"date"`
	CreatedAt   string  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   string  `json:"updated_at,omitempty" db:"updated_at"`
}

// PasswordResetToken a model for working with password reset tokens.
type PasswordResetToken struct {
	ID         int64  `json:"id" db:"id"`
	AccountID  int64  `json:"account_id" db:"account_id"`
	ResetToken string `json:"reset_token" db:"reset_token"`
	IsUsed     bool   `json:"is_used" db:"is_used"`
	ExpiresAt  string `json:"expires_at" db:"expires_at"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
