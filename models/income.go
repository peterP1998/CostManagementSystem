package models

type Income struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Value            float32   `json:"value"`
	Category         IncomeCategory     `json:"category"`
	Userid          int       `json:"userid"`
}