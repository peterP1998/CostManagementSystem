package controller


import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
)
type BalanceController struct {
	accountService service.AccountService
	balanceService service.BalanceService
	userService  service.UserService
}
func (balance *BalanceController) GetBalanceForUser(w http.ResponseWriter, r *http.Request){
	token:=balance.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=balance.accountService.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=balance.userService.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	expenses,err := service.SelectAllExpensesForUser(user.ID)
	balance.balanceService.CalculateBalanceCreateChart(w,incomes,expenses,user.ID)
}