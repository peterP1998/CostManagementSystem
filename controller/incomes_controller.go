package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
	"strconv"
)

type IncomeController struct {
	Accountervice service.AccountService
	IncomeService  service.IncomeService
	Userservice    service.UserService
}

func (incomeController IncomeController) IncomePage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/income/income.html", nil)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) GetIncomesForUser(w http.ResponseWriter, r *http.Request) {
	token := incomeController.Accountervice.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.Accountervice.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := incomeController.Userservice.SelectUserByName(username)
	utils.UserNotFound(err, w)
	incomes, err := incomeController.IncomeService.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err, w)
	err = views.CreateView(w, "static/templates/income/incomeHistory.html", incomes)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) AddIncomeForUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := incomeController.Accountervice.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.Accountervice.ParseToken(token.Value)
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	okresp := map[string]interface{}{"messg": "Income Created!"}
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	user, err := incomeController.Userservice.SelectUserByName(username)
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	err = incomeController.IncomeService.CreateIncome(user.ID, r.FormValue("description"), i, r.FormValue("category"))
	if err != nil {
		views.CreateView(w, "static/templates/income/income.html", errresp)
	}
	views.CreateView(w, "static/templates/income/income.html", okresp)
}
