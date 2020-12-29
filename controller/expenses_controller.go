package controller

import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"github.com/peterP1998/CostManagementSystem/views"
)
type ExpenseController struct {
	accountService service.AccountService
	expenseService service.ExpenseService
	userService service.UserService
}
func (expenseController ExpenseController) ExpensePage(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/expenses.html",nil)
}
/*func GetExpenesesForUser(w http.ResponseWriter, r *http.Request){
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
	expenses,err :=service.SelectAllExpensesForUser(user.ID)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expenses)
}*/
func (expenseController ExpenseController)AddExpenseForUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    token:=expenseController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=expenseController.accountService.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=expenseController.userService.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err=expenseController.expenseService.CreateExpense(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}