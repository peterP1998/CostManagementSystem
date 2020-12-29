package controller

import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"github.com/peterP1998/CostManagementSystem/views"
)
type IncomeController struct {
	accountService service.AccountService
	incomeService service.IncomeService
	userService service.UserService
}
func (incomeController IncomeController)IncomePage(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/income.html",nil)
}
func (incomeController IncomeController) GetIncomesForUser(w http.ResponseWriter, r *http.Request){
	token:=incomeController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=incomeController.accountService.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=incomeController.userService.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	views.CreateView(w,"static/templates/incomeHistory.html",incomes)
}
func (incomeController IncomeController)AddIncomeForUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    token:=incomeController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=incomeController.accountService.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=incomeController.userService.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = incomeController.incomeService.CreateIncome(user.ID,r.FormValue("description"),i)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}