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
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	user,err :=incomeController.userService.SelectUserByName(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	if err != nil {
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	views.CreateView(w,"static/templates/incomeHistory.html",incomes)
}
func (incomeController IncomeController)AddIncomeForUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    token:=incomeController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=incomeController.accountService.ParseToken(token.Value)
	if err!=nil{
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	user,err :=incomeController.userService.SelectUserByName(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	err = incomeController.incomeService.CreateIncome(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	if err!=nil{
		http.Error(w, "Something went wrong please try again.", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}