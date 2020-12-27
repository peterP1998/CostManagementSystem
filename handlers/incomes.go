package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
	"html/template"
	"strconv"
)
func IncomePage(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("static/templates/income.html")
	t.Execute(w, nil)
}
func GetIncomesForUser(w http.ResponseWriter, r *http.Request){
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
	incomes,err := service.SelectAllIncomesForUser(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(incomes)
}
func AddIncomeForUser(w http.ResponseWriter, r *http.Request){
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
	err = service.CreateIncome(user.ID,r.FormValue("description"),i)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}