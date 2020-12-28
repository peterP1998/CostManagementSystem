package service
import (
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"github.com/peterP1998/CostManagementSystem/models"
	
)
func CreateExpenseChart(userid int){
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: GetNumberOfExpensesOfOneCategory(userid,"Clothes"), Label: "Clothes"},
			{Value: GetNumberOfExpensesOfOneCategory(userid,"Rent"), Label: "Rent"},
			{Value: GetNumberOfExpensesOfOneCategory(userid,"Food"), Label: "Food"},
			{Value: GetNumberOfExpensesOfOneCategory(userid,"Bills"), Label: "Bills"},
			{Value: GetNumberOfExpensesOfOneCategory(userid,"other"), Label: "Other"},
		},
	}

	f, _ := os.Create("static/output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
func CalculateBalance(incomes []models.Income,expenses []models.Expense)(float32){
	var balance float32
	balance=0
	for _, s := range incomes {
		balance=balance+s.Value
	}
	for _, s := range expenses {
		balance=balance-s.Value
	}
	return balance
}