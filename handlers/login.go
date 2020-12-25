package handlers

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/models"
	"log"
	"net/http"
	//"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	Admin bool  `json:"admin"`
}
var jwtKey = []byte("my_secret_key")
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := models.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	res, err := db.Query("SELECT * FROM User where username=?",creds.Username)
	if err != nil {
		panic(err.Error())
	}
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
	if  user.Password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		Admin:user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}