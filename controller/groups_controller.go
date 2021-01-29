package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/utils"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
	"strconv"
)

type GroupController struct {
	Accountservice service.AccountService
	Groupservice   service.GroupService
	Userservice    service.UserService
	Expenseservice service.ExpenseService
}

func (groupController GroupController) GetCreateGroupPage(w http.ResponseWriter, r *http.Request) {
	err := views.CreateView(w, "static/templates/group/creategroup.html", nil)
	utils.InternalServerError(err, w)
}
func (groupController GroupController) GetDonateGroupPage(w http.ResponseWriter, r *http.Request) {
	groups, err := groupController.Groupservice.SelectAllGroups()
	groupmap := map[string]interface{}{
		"messg": "",
		"group": groups,
	}
	utils.InternalServerError(err, w)
	err = views.CreateView(w, "static/templates/group/donategroup.html", groupmap)
	utils.InternalServerError(err, w)
}
func (groupController GroupController) CreateGroup(w http.ResponseWriter, r *http.Request) {
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg": "Group created succesfully"}
	createGroupHtml:="static/templates/group/creategroup.html"
	r.ParseForm()
	token := groupController.Accountservice.CheckAuthBeforeOperate(r, w)
	_, admin, err := groupController.Accountservice.ParseToken(token.Value)
	if admin == false || err != nil {
		views.CreateView(w, createGroupHtml, errresp)
		return
	}
	err = groupController.Groupservice.CreateGroup(r.FormValue("money"), r.FormValue("group"))
	if err != nil {
		views.CreateView(w, createGroupHtml, errresp)
		return
	}
	views.CreateView(w,createGroupHtml, ok)
}

func (groupController GroupController) DonateMoney(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	groups, err := groupController.Groupservice.SelectAllGroups()
	errresp := map[string]interface{}{"group": groups, "messg": "Something went wrong!Try again!"}
	messg := "Group donation succesfully!"
	donateGroupHtml:="static/templates/group/donategroup.html"
	if err != nil {
		views.CreateView(w, donateGroupHtml, errresp)
		return
	}
	token := groupController.Accountservice.CheckAuthBeforeOperate(r, w)
	username, _, err := groupController.Accountservice.ParseToken(token.Value)
	if err != nil {
		views.CreateView(w, donateGroupHtml, errresp)
		return
	}
	user, err := groupController.Userservice.SelectUserByName(username)
	if err != nil {
		views.CreateView(w, donateGroupHtml, errresp)
		return
	}
	group, err := groupController.Groupservice.SelectGroupByName(r.FormValue("name"))
	transactionMoney, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		views.CreateView(w, donateGroupHtml, errresp)
		return
	}
	if group.MoneyByNow == group.TargetMoney {
		views.CreateView(w, donateGroupHtml, map[string]interface{}{"group": groups, "messg": "Target is already accomplished!"})
		return
	} else if group.MoneyByNow+float64(transactionMoney) >= group.TargetMoney {
		transactionMoney = int(group.TargetMoney - group.MoneyByNow)
		group.MoneyByNow = group.TargetMoney
		messg = messg + "Target accomplished!Well done!"
	} else {
		group.MoneyByNow = group.MoneyByNow + float64(transactionMoney)
	}
	err = groupController.Expenseservice.CreateExpense(user.ID, "Group donate", transactionMoney, "Other")
	if err != nil {
		if err.Error() == "Not enough money" {
			views.CreateView(w, donateGroupHtml, map[string]interface{}{"group": groups, "messg": "Not enough money!"})
			return
		} else {
			views.CreateView(w, donateGroupHtml, errresp)
			return
		}
	}
	err = groupController.Groupservice.UpdateGroupMoney(group.ID, int(group.MoneyByNow))
	if err != nil {
		views.CreateView(w, donateGroupHtml, errresp)
		return
	}
	ok := map[string]interface{}{"group": groups, "messg": messg}
	views.CreateView(w,donateGroupHtml, ok)
}
