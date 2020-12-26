package handlers


import (
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"encoding/json"
)

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
	token:=service.CheckAuthBeforeOperate(r,w)
	_,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	service.CreateGroupDB(w,r)
	w.WriteHeader(http.StatusCreated)
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
	db, err := db.CreateDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=service.SelectUserByName(db,username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	groupId,err:=service.SplitUrlGroup(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}*/