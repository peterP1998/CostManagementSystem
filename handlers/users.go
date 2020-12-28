package handlers

import (
	"github.com/peterP1998/CostManagementSystem/service"
	"net/http"
	"github.com/peterP1998/CostManagementSystem/models"
	"encoding/json"
	"html/template"
	"golang.org/x/crypto/bcrypt"
)
func GetUserAcc(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/user.html")
	t.Execute(w, nil)
}
func GetAdminAcc(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/admin.html")
	t.Execute(w, nil)
}
func GetCreateUserPage(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("static/templates/createuser.html")
	t.Execute(w, nil)
}
func GetDeleteUserPage(w http.ResponseWriter, r *http.Request){
	token:=service.CheckAuthBeforeOperate(r,w)
	username,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	users,err:=service.SelectAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users_without_admins:=make([]models.User, 0)
	for _, v := range users {
        if v.Name!=username&&v.Admin==false{
			users_without_admins = append(users_without_admins, v)
		}
    }
	t, _ := template.ParseFiles("static/templates/deleteuser.html")
	t.Execute(w, users_without_admins)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token:=service.CheckAuthBeforeOperate(r,w)
	tknStr := token.Value
	_,admin,err:=service.ParseToken(tknStr)
	service.CheckAdminPermission(admin,w)
	users,err:=service.SelectAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	service.CheckAuthBeforeOperate(r,w)
	userId,err:=service.SplitUrlUser(r)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user,err :=service.SelectUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	t, _ := template.ParseFiles("static/templates/createuser.html")
	token:=service.CheckAuthBeforeOperate(r,w)
	username,admin,err:=service.ParseToken(token.Value)
	if admin ==false|| err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	service.DeleteUserById(userId,w)
}
func CreateUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	t, _ := template.ParseFiles("static/templates/createuser.html")
	errresp := map[string]interface{}{"messg": "Something went wrong!Try again!"}
	ok := map[string]interface{}{"messg":"User created succesfully"}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
    if err != nil {
        t.Execute(w, errresp)
		return
	}
	admin:=false
	if r.FormValue("admin")=="yes"{
       admin=true
	}
	service.CreateUserDB(w,r.FormValue("username"),r.FormValue("email"),string(hashedPassword),admin)
	t.Execute(w, ok)
}
func RegisterUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.RegisterUserDB(w,r.FormValue("username"),r.FormValue("email"),string(hashedPassword))
	http.Redirect(w, r, "/api/login", http.StatusSeeOther)
}
