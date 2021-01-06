package controller


import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/views"
	"github.com/peterP1998/CostManagementSystem/utils"
	"strconv"
)
type GroupController struct {
	accountService service.AccountService
	groupService service.GroupService
	userService service.UserService
	expenseService service.ExpenseService
}
func (groupController GroupController) GetCreateGroupPage(w http.ResponseWriter, r *http.Request){
	err:=views.CreateView(w,"static/templates/creategroup.html",nil)
    utils.InternalServerError(err,w)
}
func (groupController GroupController) GetDonateGroupPage(w http.ResponseWriter, r *http.Request){
	groups,err:=groupController.groupService.SelectAllGroups()
	groupmap:=map[string]interface{}{
		"messg":"",
		"group":groups,
	}
	utils.InternalServerError(err,w)
	err=views.CreateView(w,"static/templates/donategroup.html",groupmap)
    utils.InternalServerError(err,w)
}
func (groupController GroupController) CreateGroup(w http.ResponseWriter, r *http.Request){
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg":"Group created succesfully"}
	r.ParseForm()
	token:=groupController.accountService.CheckAuthBeforeOperate(r,w)
	_,admin,err:=groupController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		views.CreateView(w,"static/templates/creategroup.html",errresp)
	}
	err=groupController.groupService.CreateGroup(r.FormValue("money"),r.FormValue("group"))
	if err!=nil{
		views.CreateView(w,"static/templates/creategroup.html",errresp)
	}
	views.CreateView(w,"static/templates/creategroup.html",ok)
}

func (groupController GroupController) DonateMoney(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	groups,err:=groupController.groupService.SelectAllGroups()
	errresp := map[string]interface{}{"group":groups,"messg": "Something went wrong!Try again!"}
	messg:="Group donation succesfully"
	if err != nil {
		views.CreateView(w,"static/templates/donategroup.html",errresp)
		return
	}
	token:=groupController.accountService.CheckAuthBeforeOperate(r,w)
	username,_,err:=groupController.accountService.ParseToken(token.Value)
	if err!=nil{
		views.CreateView(w,"static/templates/donategroup.html",errresp)
		return
	}
	user,err :=groupController.userService.SelectUserByName(username)
	if err != nil {
		views.CreateView(w,"static/templates/donategroup.html",errresp)
		return
	}
	group,err:=groupController.groupService.SelectGroupByName(r.FormValue("name"))
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		views.CreateView(w,"static/templates/donategroup.html",errresp)
	    return
	}
	if group.MoneyByNow == group.TargetMoney{
		views.CreateView(w,"static/templates/donategroup.html", map[string]interface{}{"group":groups,"messg": "Target is already accomplished!"})
		return
	}else if group.MoneyByNow+float64(i)>=group.TargetMoney{
		i=int(group.TargetMoney-group.MoneyByNow)
		group.MoneyByNow=group.TargetMoney
		messg=messg+"Target accomplished!Well done!"
	}else{
        group.MoneyByNow=group.MoneyByNow+float64(i)
	}
	err=groupController.groupService.UpdateGroupMoney(group.ID,int(group.MoneyByNow))
	if err != nil {
		views.CreateView(w,"static/templates/donategroup.html",errresp)
	    return
	}
	groupController.expenseService.CreateExpense(user.ID,"Group donate",i,"Other")
	ok := map[string]interface{}{"group":groups,"messg":messg}
	views.CreateView(w,"static/templates/donategroup.html",ok)
}