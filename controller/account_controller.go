package controller

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"github.com/peterP1998/CostManagementSystem/views"
	"net/http"
)

type AccountController struct {
    account service.AccountService
}
func (controller AccountController) GetLoginForm(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/accounts/index.html",nil)
}
func (controller AccountController)Welcome(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/welcome.html",nil)
}
func (controller AccountController) GetRegister(w http.ResponseWriter, r *http.Request){
	views.CreateView(w,"static/templates/accounts/signup.html",nil)
}
func (controller AccountController) Account(w http.ResponseWriter, r *http.Request){
	token:=controller.account.CheckAuthBeforeOperate(r,w)
	_,admin,err:=controller.account.ParseToken(token.Value)
	if err!=nil{
        http.Redirect(w, r, "/api/login", http.StatusSeeOther)
	}
	if admin ==false{
		views.CreateView(w,"static/templates/accounts/user.html",nil)
	}else {
		views.CreateView(w,"static/templates/accounts/admin.html",nil)
	}
}
func (controller AccountController) Signin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	controller.account.Login(r.FormValue("password"),r.FormValue("username"),w)
	http.Redirect(w, r, "/api/account", http.StatusSeeOther)
}
func (controller AccountController) Logout(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

