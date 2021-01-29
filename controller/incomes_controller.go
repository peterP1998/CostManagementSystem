package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
	"strconv"
)

type IncomeController struct {
	Accountservice service.AccountService
	Incomeservice  service.IncomeService
	Userservice    service.UserService
}

func (incomeController IncomeController) IncomePage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/income/income.html", nil)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) GetIncomesForUser(w http.ResponseWriter, r *http.Request) {
	token := incomeController.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.Accountservice.ParseToken(token.Value)
	utils.InternalServerError(err, w)
	user, err := incomeController.Userservice.SelectUserByName(username)
	utils.UserNotFound(err, w)
	incomes, err := incomeController.Incomeservice.SelectAllIncomesForUser(user.ID)
	utils.InternalServerError(err, w)
	err = views.CreateView(w, "static/templates/income/incomeHistory.html", incomes)
	utils.InternalServerError(err, w)
}
func (incomeController IncomeController) AddIncomeForUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := incomeController.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := incomeController.Accountservice.ParseToken(token.Value)
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	okresp := map[string]interface{}{"messg": "Income Created!"}
	incomesHtml:="static/templates/income/income.html"
	if err != nil {
		views.CreateView(w,incomesHtml, errresp)
	}
	user, err := incomeController.Userservice.SelectUserByName(username)
	if err != nil {
		views.CreateView(w,incomesHtml, errresp)
	}
	i, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		views.CreateView(w,incomesHtml, errresp)
	}
	err = incomeController.Incomeservice.CreateIncome(user.ID, r.FormValue("description"), i, r.FormValue("category"))
	if err != nil {
		views.CreateView(w, incomesHtml, errresp)
	}
	views.CreateView(w,incomesHtml, okresp)
}
