package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
)

func GetExpenesesForUser(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=service.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	expenses,err :=service.SelectAllExpensesForUser(user.ID)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expenses)
}
func AddExpenseForUser(w http.ResponseWriter, r *http.Request){
    token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user,err :=service.SelectUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err=service.CreateExpense(user.ID,r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}