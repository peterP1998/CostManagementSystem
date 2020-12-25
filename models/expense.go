package models

type Expense struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Value            float32     `json:"value"`
	Userid          int    `json:"userid"`
}