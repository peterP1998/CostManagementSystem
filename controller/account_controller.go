package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"html/template"
	"golang.org/x/crypto/bcrypt"
)


func GetLoginForm(w http.ResponseWriter, r *http.Request){
	CreateView(w,"static/templates/index.html",nil)
}
func Welcome(w http.ResponseWriter, r *http.Request){
	CreateView(w,"static/templates/welcome.html",nil)
}
func GetRegister(w http.ResponseWriter, r *http.Request){
	CreateView(w,"static/templates/signup.html",nil)
}
func Account(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	_,admin,err:=service.ParseToken(token.Value)
	if err!=nil{
        http.Redirect(w, r, "/api/login", http.StatusSeeOther)
	}
	if admin ==false{
		CreateView(w,"static/templates/user.html",nil)
	}else {
		CreateView(w,"static/templates/admin.html",nil)
	}
}
func Signin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	user,err :=service.SelectUserByName(r.FormValue("username"))
	if err != nil {
		CreateView(w,"static/templates/index.html",errresp)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),  []byte(r.FormValue("password")))
	if  err!=nil {
		CreateView(w,"static/templates/index.html",errresp)
		return
	}
	err=service.CreateAndConfigureToken(user,w)
	if err != nil {
		CreateView(w,"static/templates/index.html",errresp)
		return
	}
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}
func Logout(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
