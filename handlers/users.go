package handlers

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"encoding/json"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token:=service.CheckAuthBeforeOperate(r,w)
	tknStr := token.Value
	_,admin,err:=service.ParseToken(tknStr)
	service.CheckAdminPermission(admin,w)
	users,err:=service.SelectAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrlUser(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user,err :=service.SelectUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrlUser(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	service.DeleteUserById(userId,w)
}
func CreateUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	service.CreateUserDB(w,r.FormValue("username"),r.FormValue("email"),r.FormValue("password"))
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}
