package models

type Expense struct {
	ID          int    `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`
	Category    string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
	Amount      int64  `json:"amount" db:"amount"`
	Date        string `json:"date" db:"date"`
}
