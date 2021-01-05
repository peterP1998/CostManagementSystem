package models

type Income struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Value            float32   `json:"value"`
	Category         string     `json:"category"`
	Userid          int       `json:"userid"`
}