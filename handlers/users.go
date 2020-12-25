package handlers

import (
	"github.com/peterP1998/CostManagementSystem/db"
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
	db, err := db.CreateDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	users,err:=service.SelectAllUsers(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrl(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db, err := db.CreateDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=service.SelectUserById(db,userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer db.Close()
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrl(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db, err := db.CreateDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.DeleteUserById(db,userId,w)
}
func CreateUser(w http.ResponseWriter, r *http.Request){
	db, err := db.CreateDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.CreateUserDB(db,w,r)
	w.WriteHeader(http.StatusCreated)
}
