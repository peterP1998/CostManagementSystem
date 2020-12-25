package service

import (
	"encoding/json"
	"github.com/peterP1998/CostManagementSystem/models"
	"io"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)
var jwtKey = []byte("my_secret_key")
func DecodeJsonCredentials(body io.Reader)(models.Credentials,error){
    var creds models.Credentials
	err := json.NewDecoder(body).Decode(&creds)
	if err != nil {
		return models.Credentials{},err
	}
	return creds,nil
}
func ParseToken(tokenString string) (string,bool,error){
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"),nil
	})
	if err!=nil{
		return "",false,err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Username,claims.Admin,nil
	} else {
		return "",false,err
	}
}
func CreateAndConfigureToken(user models.User,w http.ResponseWriter)(error){
	claims,expirationTime:=configureToken(user)
	tokenString, err := createToken(claims)
	if err != nil {
		return err
	}
	setCookieToken(tokenString,expirationTime,w)
	return nil
}
func configureToken(user models.User)(*models.Claims,time.Time){
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: user.Name,
		Admin:user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	return claims,expirationTime
}
func createToken(claims *models.Claims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString,err
}
func setCookieToken(tokenString string,expirationTime time.Time,w http.ResponseWriter){
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func checkForAuthorization(r *http.Request)(*http.Cookie,error){
	toekn, err := r.Cookie("token")
	if err != nil {
		return nil,err
	}
	return toekn,nil
}
func authorizationResponses(w http.ResponseWriter,err error){
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func CheckAuthBeforeOperate(r *http.Request,w http.ResponseWriter)(*http.Cookie){
	token, err := checkForAuthorization(r)
	authorizationResponses(w,err)
	return token
}
func CheckAdminPermission(admin bool,w http.ResponseWriter){
    if admin ==false{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}