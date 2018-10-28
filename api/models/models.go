package models

type Expense struct {
	ID          int    `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`
	Category    string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
	Amount      int64  `json:"amount" db:"amount"`
	Date        string `json:"date" db:"date"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type Type struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type Category struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
