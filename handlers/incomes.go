package handlers

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/db"
	"github.com/peterP1998/CostManagementSystem/models"
	"log"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/service"
)

func GetIncomesForUser(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
    db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	user,err :=service.SelectUserByName(db,username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res,err :=db.Query("select * from Income where userid=?;",user.ID)
	incomes := make([]models.Income, 0)
	for res.Next() {
		var income models.Income
		res.Scan(&income.ID, &income.Description, &income.Value, &income.Userid)
		incomes = append(incomes, income)
	}
	json.NewEncoder(w).Encode(incomes)
}
func AddIncomeForUser(w http.ResponseWriter, r *http.Request){
    token:=service.CheckAuthBeforeOperate(r,w)
	username,_,err:=service.ParseToken(token.Value)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
    db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	user,err :=service.SelectUserByName(db,username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var income models.Income
	err = json.NewDecoder(r.Body).Decode(&income)
	_,err =db.Query("insert into Income(description,value,userid) Values(?,?,?);",income.Description,income.Value,user.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}