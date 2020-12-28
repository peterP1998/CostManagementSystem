package handlers


import (
	"html/template"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	expenses,err := service.SelectAllExpensesForUser(user.ID)
	balance:=service.CalculateBalance(incomes,expenses)
	t, _ := template.ParseFiles("static/templates/balance.html")
	myvar := map[string]interface{}{"Balance": balance}
	service.CreateExpenseChart(user.ID)
	t.Execute(w, myvar)
}