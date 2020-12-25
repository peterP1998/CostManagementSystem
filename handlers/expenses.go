package handlers

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/db"
	"github.com/peterP1998/CostManagementSystem/models"
	"log"
	"net/http"
	"fmt"
)

func GetExpenesesForUser(w http.ResponseWriter, r *http.Request){
    toekn, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := toekn.Value
	username,_,err:=ParseToken(tknStr)
	fmt.Println(username)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
    db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	err =db.QueryRow("select * from User where username=?;",username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		log.Fatal(err)
		return
	}
	res,err :=db.Query("select * from Expense where userid=?;",user.ID)
	expenses := make([]models.Expense, 0)
	for res.Next() {
		var expense models.Expense
		res.Scan(&expense.ID, &expense.Description, &expense.Value, &expense.Userid)
		expenses = append(expenses, expense)
	}
	json.NewEncoder(w).Encode(expenses)
}
func AddExpenseForUser(w http.ResponseWriter, r *http.Request){
    toekn, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := toekn.Value
	username,_,err:=ParseToken(tknStr)
	var user models.User
	db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err =db.QueryRow("select * from User where username=?;",username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		log.Fatal(err)
		return
	}
	var expense models.Expense
	err = json.NewDecoder(r.Body).Decode(&expense)
	_,err =db.Query("insert into Expense(description,value,userid) Values(?,?,?);",expense.Description,expense.Value,user.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
}