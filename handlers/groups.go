package handlers


import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"encoding/json"
	"html/template"
	"strconv"
)
func GetCreateGroupPage(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("static/templates/creategroup.html")
	t.Execute(w, nil)
}
func GetGroup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    service.CheckAuthBeforeOperate(r,w)
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
}
func CreateGroup(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/creategroup.html")
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg":"Group created succesfully"}
	r.ParseForm()
	token:=service.CheckAuthBeforeOperate(r,w)
	_,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		t.Execute(w, errresp)
		return
	}
	i, err := strconv.Atoi(r.FormValue("money"))
	if err!=nil{
		t.Execute(w, errresp)
		return
	}
	service.CreateGroupDB(i,r.FormValue("group"))
	t.Execute(w, ok)
}
func DeleteGroup(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	_,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	groupId,err:=service.SplitUrlGroup(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.DeleteGroup(groupId,w)
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