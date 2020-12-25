package handlers

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/models"
	"log"
	"net/http"
	//"fmt"
	"strings"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	db, err := models.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	res, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	users := make([]models.User, 0)
	for res.Next() {
		var user models.User
		res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	p := strings.Split(r.URL.Path, "/user/")
	userId,err:=strconv.Atoi(p[len(p)-1])
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db, err := models.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	res, err :=db.Query("select * from User where id=?;",userId)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer res.Close()
	var user models.User
	count := 0
    for res.Next() { 
	   res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
       count += 1 
	}
	if count!=1{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
