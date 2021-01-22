package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"net/http"
)

type BalanceController struct {
	accountService service.AccountService
	balanceService service.BalanceService
	userService    service.UserService
}

func (balance *BalanceController) GetBalanceForUser(w http.ResponseWriter, r *http.Request) {
	token := balance.accountService.CheckAuthBeforeOperate(r, w)
	username, _, err := balance.accountService.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := balance.userService.SelectUserByName(username)
	utils.InternalServerError(err, w)
	incomes, err := service.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err, w)
	expenses, err := service.SelectAllExpensesForUser(user.ID)
	utils.InternalServerError(err, w)
	balance.balanceService.CalculateBalanceCreateChart(w, incomes, expenses, user.ID)
}
