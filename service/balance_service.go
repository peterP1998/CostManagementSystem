package service

import (
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/wcharczuk/go-chart/v2"
	"net/http"
	"os"
)

type BalanceService struct {
}


func (balanceService BalanceService) CalculateBalanceCreateChart(w http.ResponseWriter, incomes []models.Income, expenses []models.Expense, userid int) {
	balance := CalculateBalance(incomes, expenses)
	createExpenseChart(userid)
	createIncomeChart(userid)
	views.CreateView(w, "static/templates/balance.html", map[string]interface{}{"Balance": balance})
}
func createExpenseChart(userid int) {
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: getValueOfExpensesOfOneCategory(userid, "Clothes"), Label: "Clothes"},
			{Value: getValueOfExpensesOfOneCategory(userid, "Rent"), Label: "Rent"},
			{Value: getValueOfExpensesOfOneCategory(userid, "Food"), Label: "Food"},
			{Value: getValueOfExpensesOfOneCategory(userid, "Bills"), Label: "Bills"},
			{Value: getValueOfExpensesOfOneCategory(userid, "other"), Label: "Other"},
		},
	}

	f, _ := os.Create("static/output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
func createIncomeChart(userid int) {
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: getValueOfIncomesOfOneCategory(userid, "Salary"), Label: "Salary"},
			{Value: getValueOfIncomesOfOneCategory(userid, "Gift"), Label: "Gift"},
			{Value: getValueOfIncomesOfOneCategory(userid, "Found"), Label: "Found"},
			{Value: getValueOfIncomesOfOneCategory(userid, "Sell"), Label: "Sell"},
		},
	}

	f, _ := os.Create("static/income.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
func CalculateBalance(incomes []models.Income, expenses []models.Expense) float32 {
	var balance float32
	balance = 0
	for _, s := range incomes {
		balance = balance + s.Value
	}
	for _, s := range expenses {
		balance = balance - s.Value
	}
	return balance
}
func getValueOfExpensesOfOneCategory(id int, category string) float64 {
	var cnt float64
	res, _ := models.DB.Query(`select * from Expense where userid=? and category=?;`, id, category)
	cnt = 0
	for res.Next() {
		var expense models.Expense
		res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Category, &expense.Userid)
		cnt = cnt + float64(expense.Value)
	}
	return cnt
}
func getValueOfIncomesOfOneCategory(id int, category string) float64 {
	var cnt float64
	res, _ := models.DB.Query(`select * from Income where userid=? and category=?;`, id, category)
	cnt = 0
	for res.Next() {
		var income models.Income
		res.Scan(&income.ID, &income.Description, &income.Value, &income.Category, &income.Userid)
		cnt = cnt + float64(income.Value)
	}
	return cnt
}
