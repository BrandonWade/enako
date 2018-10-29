package models

type Expense struct {
	ID          int64  `json:"id" db:"id"`
	UserID      int64  `json:"user_id,omitempty" db:"user_id"`
	Type        string `json:"type,omitempty" db:"type"`
	Category    string `json:"category,omitempty" db:"category"`
	Description string `json:"description,omitempty" db:"description"`
	Amount      int64  `json:"amount,omitempty" db:"amount"`
	Date        string `json:"date,omitempty" db:"date"`
	CreatedAt   string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty" db:"updated_at"`
}

type Type struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name,omitempty" db:"name"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

type Category struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
