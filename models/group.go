package models

type Group struct {
	ID             int       `json:"id"`
	GroupName      string    `json:"groupname"`
	MoneyByNow     float64    `json:"moneybynow"`
	TargetMoney    float64    `json:"targetmoney"`
}