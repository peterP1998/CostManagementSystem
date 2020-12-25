package handlers

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/db"
	"log"
	"net/http"
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

	db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var user models.User
	err =db.QueryRow("SELECT * FROM User where username=?",creds.Username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
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
func Logout(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)

	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
func ParseToken(tokenString string) (string,bool,error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"),nil
	})
	if err!=nil{
		return "",false,err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username,claims.Admin,nil
	} else {
		return "",false,err
	}
}
