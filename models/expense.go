package models

type Expense struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Value            float32     `json:"value"`
	Category         ExpenseCategory     `json:"category"`
	Userid          int    `json:"userid"`
}