package controller

import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/peterP1998/CostManagementSystem/utils"
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
	utils.InternalServerError(err,w)
	user,err :=incomeController.userService.SelectUserByName(username)
	utils.UserNotFound(err,w)
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err,w)
	views.CreateView(w,"static/templates/incomeHistory.html",incomes)
}
func (incomeController IncomeController)AddIncomeForUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    token:=incomeController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=incomeController.accountService.ParseToken(token.Value)
	utils.InternalServerError(err,w)
	user,err :=incomeController.userService.SelectUserByName(username)
	utils.UserNotFound(err,w)
	i, err := strconv.Atoi(r.FormValue("value"))
	utils.InternalServerError(err,w)
	err = incomeController.incomeService.CreateIncome(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	utils.InternalServerError(err,w)
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}