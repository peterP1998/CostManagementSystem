package service

import (
	"fmt"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/repository"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/wcharczuk/go-chart/v2"
	"net/http"
	"os"
)

type BalanceService struct {
}

func (balanceService BalanceService) CalculateBalanceCreateChart(w http.ResponseWriter, incomes []models.Income, expenses []models.Expense, userid int) {
	balance := CalculateBalance(incomes, expenses)
	createChart("expense"+fmt.Sprint(userid), userid, createArrayOfExepnses(userid))
	createChart("income"+fmt.Sprint(userid), userid, createArrayOfIncomes(userid))
	views.CreateView(w, "static/templates/balance/balance.html", map[string]interface{}{"Balance": balance, "Income": "income" + fmt.Sprint(userid), "Expense": "expense" + fmt.Sprint(userid)})
}
func createArrayOfExepnses(userid int) []chart.Value {
	var expenseRepo = repository.ExpenseRepository{}
	values := []chart.Value{
		{Value: getValueOfExpensesOfOneCategory(userid, "Clothes", expenseRepo), Label: "Clothes"},
		{Value: getValueOfExpensesOfOneCategory(userid, "Rent", expenseRepo), Label: "Rent"},
		{Value: getValueOfExpensesOfOneCategory(userid, "Food", expenseRepo), Label: "Food"},
		{Value: getValueOfExpensesOfOneCategory(userid, "Bills", expenseRepo), Label: "Bills"},
		{Value: getValueOfExpensesOfOneCategory(userid, "Other", expenseRepo), Label: "Other"},
	}
	return values
}
func createArrayOfIncomes(userid int) []chart.Value {
	var incomeRepo = repository.IncomeRepository{}
	values := []chart.Value{
		{Value: getValueOfIncomesOfOneCategory(userid, "Salary", incomeRepo), Label: "Salary"},
		{Value: getValueOfIncomesOfOneCategory(userid, "Gift", incomeRepo), Label: "Gift"},
		{Value: getValueOfIncomesOfOneCategory(userid, "Found", incomeRepo), Label: "Found"},
		{Value: getValueOfIncomesOfOneCategory(userid, "Sell", incomeRepo), Label: "Sell"},
	}
	return values
}
func createChart(pictureName string, userid int, values []chart.Value) {
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: values,
	}

	f, _ := os.Create("static/" + pictureName)
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
func getValueOfExpensesOfOneCategory(id int, category string, expenseRepo ExpenseRepositoryInterface) float64 {
	var cnt float64
	res, _ := expenseRepo.GetExpensesByCategoryAndUserId(id, category)
	cnt = 0.0
	if res != nil {
		for _, v := range res {
			cnt = cnt + float64(v.Value)
		}
	}
	return cnt
}
func getValueOfIncomesOfOneCategory(id int, category string, incomeRepo IncomeRepositoryInterface) float64 {
	var cnt float64
	res, _ := incomeRepo.GetIncomesByCategoryAndUserId(id, category)
	cnt = 0.0
	if res != nil {
		for _, v := range res {
			cnt = cnt + float64(v.Value)
		}
	}
	return cnt
}
