package service
import (
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
)
type BalanceService struct {

}
func createExpenseChart(userid int){
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: getNumberOfExpensesOfOneCategory(userid,"Clothes"), Label: "Clothes"},
			{Value: getNumberOfExpensesOfOneCategory(userid,"Rent"), Label: "Rent"},
			{Value: getNumberOfExpensesOfOneCategory(userid,"Food"), Label: "Food"},
			{Value: getNumberOfExpensesOfOneCategory(userid,"Bills"), Label: "Bills"},
			{Value: getNumberOfExpensesOfOneCategory(userid,"other"), Label: "Other"},
		},
	}

	f, _ := os.Create("static/output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
func createIncomeChart(userid int){
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: getNumberOfIncomesOfOneCategory(userid,"Salary"), Label: "Salary"},
			{Value: getNumberOfIncomesOfOneCategory(userid,"Gift"), Label: "Gift"},
			{Value: getNumberOfIncomesOfOneCategory(userid,"Found"), Label: "Found"},
			{Value: getNumberOfIncomesOfOneCategory(userid,"Sell"), Label: "Sell"},
		},
	}

	f, _ := os.Create("static/income.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
func calculateBalance(incomes []models.Income,expenses []models.Expense)(float32){
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
func (balanceService BalanceService) CalculateBalanceCreateChart(w http.ResponseWriter,incomes []models.Income,expenses []models.Expense,userid int){
	balance:=calculateBalance(incomes,expenses)
	createExpenseChart(userid)
	createIncomeChart(userid)
	views.CreateView(w,"static/templates/balance.html",map[string]interface{}{"Balance": balance})
}
func  getNumberOfExpensesOfOneCategory(id int,category string)(float64){
	var cnt float64
	res,err:= models.DB.Query(`select * from Expense where userid=? and category=?;`,id,category)
	if err!=nil{
		
	}
	cnt=0
	for res.Next() {
		var expense models.Expense
		res.Scan(&expense.ID, &expense.Description, &expense.Value,&expense.Category, &expense.Userid)
		cnt=cnt+float64(expense.Value)
	}
    return cnt
}
func  getNumberOfIncomesOfOneCategory(id int,category string)(float64){
	var cnt float64
	res,err:= models.DB.Query(`select * from Income where userid=? and category=?;`,id,category)
	if err!=nil{
		
	}
	cnt=0
	for res.Next() {
		var income models.Income
		res.Scan(&income.ID, &income.Description, &income.Value,&income.Category, &income.Userid)
		cnt=cnt+float64(income.Value)
	}
    return cnt
}