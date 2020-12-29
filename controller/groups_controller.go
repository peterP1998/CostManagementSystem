package controller


import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/views"
)
type GroupController struct {
	accountService service.AccountService
	groupService service.GroupService
}
func (groupController GroupController) GetCreateGroupPage(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/creategroup.html",nil)
}
/*func (groupController GroupController) GetGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    groupController.accountService.CheckAuthBeforeOperate(r,w)
	groupId,err:=service.SplitUrlGroup(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	group,err:=service.SelectGroupById(groupId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(group)
}*/
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
func (groupController GroupController) DeleteGroup(w http.ResponseWriter, r *http.Request){
	token:=groupController.accountService.CheckAuthBeforeOperate(r,w)
	_,admin,err:=groupController.accountService.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	groupId,err:=service.SplitUrlGroup(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	groupController.groupService.DeleteGroup(groupId,w)
}
/*func DonateMoney(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//user,err :=service.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	groupId,err:=service.SplitUrlGroup(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
//	group,err:=service.SelectGroupById(groupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}*/