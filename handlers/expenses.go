package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"strconv"
	"html/template"
)
func ExpensePage(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("static/templates/expenses.html")
	t.Execute(w, nil)
}
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
	r.ParseForm()
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
	i, err := strconv.Atoi(r.FormValue("value"))
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err=service.CreateExpense(user.ID,r.FormValue("description"),i,r.FormValue("category"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}