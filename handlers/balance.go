package handlers


import (
	"html/template"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/wcharczuk/go-chart/v2"
	"fmt"
	"os"
	
)
func GetBalanceForUser(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=service.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	expenses,err := service.SelectAllExpensesForUser(user.ID)
	var balance float32
	balance=0
	for _, s := range incomes {
		balance=balance+s.Value
	}
	for _, s := range expenses {
		balance=balance-s.Value
	}
	t, _ := template.ParseFiles("static/templates/balance.html")
	myvar := map[string]interface{}{"Balance": balance}
	fmt.Println(service.GetNumberOfExpensesOfOneCategory(user.ID,"Clothes"))
	pie := chart.PieChart{
		Width:  256,
		Height: 256,
		Values: []chart.Value{
			{Value: service.GetNumberOfExpensesOfOneCategory(user.ID,"Clothes"), Label: "Clothes"},
			{Value: service.GetNumberOfExpensesOfOneCategory(user.ID,"Rent"), Label: "Rent"},
			{Value: service.GetNumberOfExpensesOfOneCategory(user.ID,"Food"), Label: "Food"},
			{Value: service.GetNumberOfExpensesOfOneCategory(user.ID,"Bills"), Label: "Bills"},
			{Value: service.GetNumberOfExpensesOfOneCategory(user.ID,"other"), Label: "Other"},
		},
	}

	f, _ := os.Create("static/output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
	t.Execute(w, myvar)
}