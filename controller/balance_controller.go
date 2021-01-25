package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"net/http"
)

type BalanceController struct {
	Accountservice service.AccountService
	Balanceservice service.BalanceService
	Userservice    service.UserService
	Expenseservice service.ExpenseService
	Incomeservice  service.IncomeService
}

func (balance *BalanceController) GetBalanceForUser(w http.ResponseWriter, r *http.Request) {
	token := balance.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := balance.Accountservice.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := balance.Userservice.SelectUserByName(username)
	utils.InternalServerError(err, w)
	incomes, err := balance.Incomeservice.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err, w)
	expenses, err := balance.Expenseservice.SelectAllExpensesForUser(user.ID)
	utils.InternalServerError(err, w)
	balance.Balanceservice.CalculateBalanceCreateChart(w, incomes, expenses, user.ID)
}
