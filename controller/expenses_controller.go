package controller

import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/peterP1998/CostManagementSystem/utils"
)
type ExpenseController struct {
	accountService service.AccountService
	expenseService service.ExpenseService
	userService service.UserService
}
func (expenseController ExpenseController) ExpensePage(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/expenses.html",nil)

}
func (expenseController ExpenseController) GetExpenesesForUser(w http.ResponseWriter, r *http.Request){
	token:=expenseController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=expenseController.accountService.ParseToken(token.Value)
	utils.InternalServerError(err,w)
	user,err :=expenseController.userService.SelectUserByName(username)
	utils.UserNotFound(err,w)
	expenses,err :=service.SelectAllExpensesForUser(user.ID)
	utils.InternalServerError(err,w)
	views.CreateView(w,"static/templates/expenseHistory.html",expenses)
}
func (expenseController ExpenseController)AddExpenseForUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    token:=expenseController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=expenseController.accountService.ParseToken(token.Value)
	utils.InternalServerError(err,w)
	user,err :=expenseController.userService.SelectUserByName(username)
	utils.UserNotFound(err,w)
	i, err := strconv.Atoi(r.FormValue("value"))
	utils.InternalServerError(err,w)
	err=expenseController.expenseService.CreateExpense(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	utils.InternalServerError(err,w)
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}