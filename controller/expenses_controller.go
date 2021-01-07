package controller

import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/peterP1998/CostManagementSystem/utils"
	///"errors"
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
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	okresp:= map[string]interface{}{"messg": "Expense Created!"}
    token:=expenseController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=expenseController.accountService.ParseToken(token.Value)
	if err!=nil{
		views.CreateView(w,"static/templates/expenses.html",errresp)
	}
	user,err :=expenseController.userService.SelectUserByName(username)
	if err!=nil{
		views.CreateView(w,"static/templates/expenses.html",errresp)
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		views.CreateView(w,"static/templates/expenses.html",errresp)
	}
	err=expenseController.expenseService.CreateExpense(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	if err!=nil{
	    if err.Error()=="Not enough money"{
		    views.CreateView(w,"static/templates/expenses.html",map[string]interface{}{"messg": "Not enough money"})
			return
		}else{
			views.CreateView(w,"static/templates/expenses.html",errresp)
			return 
		}
	}
	views.CreateView(w,"static/templates/expenses.html",okresp)
}