package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	Accountservice service.AccountService
	Expenseservice service.ExpenseService
	Userservice    service.UserService
}

func (expenseController ExpenseController) ExpensePage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/expense/expenses.html", nil)
	utils.InternalServerError(err, w)
}
func (expenseController ExpenseController) GetExpenesesForUser(w http.ResponseWriter, r *http.Request) {
	token := expenseController.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := expenseController.Accountservice.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := expenseController.Userservice.SelectUserByName(username)
	utils.UserNotFound(err, w)
	expenses, err := expenseController.Expenseservice.SelectAllExpensesForUser(user.ID)
	utils.InternalServerError(err, w)
	err = views.CreateView(w, "static/templates/expense/expenseHistory.html", expenses)
	utils.InternalServerError(err, w)
}
func (expenseController ExpenseController) AddExpenseForUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	okresp := map[string]interface{}{"messg": "Expense Created!"}
	expenseHtml:= "static/templates/expense/expenses.html"
	token := expenseController.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := expenseController.Accountservice.ParseToken(token.Value)
	if err != nil {
		views.CreateView(w, expenseHtml, errresp)
	}
	user, err := expenseController.Userservice.SelectUserByName(username)
	if err != nil {
		views.CreateView(w,expenseHtml, errresp)
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		views.CreateView(w,expenseHtml, errresp)
	}
	err = expenseController.Expenseservice.CreateExpense(user.ID, r.FormValue("description"), i, r.FormValue("category"))
	if err != nil {
		if err.Error() == "Not enough money" {
			views.CreateView(w, expenseHtml, map[string]interface{}{"messg": "Not enough money"})
			return
		} else {
			views.CreateView(w,expenseHtml, errresp)
			return
		}
	}
	views.CreateView(w,expenseHtml, okresp)
}
