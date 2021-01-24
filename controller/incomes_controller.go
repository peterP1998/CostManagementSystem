package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
	"strconv"
)

type IncomeController struct {
	accountService service.AccountService
	incomeService  service.IncomeService
	userService    service.UserService
}

func (incomeController IncomeController) IncomePage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/income/income.html", nil)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) GetIncomesForUser(w http.ResponseWriter, r *http.Request) {
	token := incomeController.accountService.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.accountService.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := incomeController.userService.SelectUserByName(username)
	utils.UserNotFound(err, w)
	incomes, err := service.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err, w)
	err = views.CreateView(w, "static/templates/income/incomeHistory.html", incomes)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) AddIncomeForUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := incomeController.accountService.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.accountService.ParseToken(token.Value)
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	okresp := map[string]interface{}{"messg": "Income Created!"}
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	user, err := incomeController.userService.SelectUserByName(username)
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	err = incomeController.incomeService.CreateIncome(user.ID, r.FormValue("description"), i, r.FormValue("category"))
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	views.CreateView(w, "static/templates/income/income.html", okresp)
}
