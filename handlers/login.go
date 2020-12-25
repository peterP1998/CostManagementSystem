package handlers

import (
	"github.com/peterP1998/CostManagementSystem/db"
	"github.com/peterP1998/CostManagementSystem/service"
	"log"
	"net/http"
)

var jwtKey = []byte("my_secret_key")
func Signin(w http.ResponseWriter, r *http.Request) {
	creds,err := service.DecodeJsonCredentials(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	user,err :=service.SelectUserByName(db,creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if  user.Password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err=service.CreateAndConfigureToken(user,w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func Logout(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)

	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}

